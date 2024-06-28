package postgres

import (
	"database/sql"
	"sca/.gen/jet/sca/public/model"
	. "sca/.gen/jet/sca/public/table"
	"sca/internal/context"
	domain "sca/internal/domain/cat"
	derr "sca/internal/error/data"
	abstract "sca/internal/repository/cat"
	commJet "sca/internal/repository/common/postgres_jet"
	"sca/internal/util"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

type catPostgresRepo struct {
	conn *sql.DB
}

func NewCatPostgresRepo(conn *sql.DB) abstract.CatRepo {
	return &catPostgresRepo{
		conn,
	}
}

func (r *catPostgresRepo) CreateCat(
	ctx *context.Context,
	create *domain.CreateCat,
) (*domain.Cat, derr.DataError) {
	modCat := newModelFromCreate(create)
	if err := Cats.INSERT(Cats.Name, Cats.YearsOfExperience, Cats.Breed, Cats.SalaryInCents).
		MODEL(modCat).
		RETURNING(Cats.AllColumns).QueryContext(ctx, r.conn, &modCat); err != nil {
		return nil, r.newErr(err)
	}

	return newDomainFromModel(&modCat), nil
}

func (r *catPostgresRepo) UpdateCat(
	ctx *context.Context,
	update *domain.UpdateCat,
) (*domain.Cat, derr.DataError) {
	modCat := newModelFromUpdate(update)
	if err := Cats.UPDATE(Cats.SalaryInCents).
		MODEL(modCat).
		WHERE(Cats.ID.EQ(UUID(update.Id))).
		RETURNING(Cats.AllColumns).QueryContext(ctx, r.conn, &modCat); err != nil {
		return nil, r.newErr(err, derr.NotFound)
	}

	return newDomainFromModel(&modCat), nil
}

func (r *catPostgresRepo) GetCat(
	ctx *context.Context,
	sel uuid.UUID,
) (*domain.Cat, derr.DataError) {
	var modCat model.Cats
	if err := Cats.SELECT(Cats.AllColumns).
		FROM(Cats).
		WHERE(Cats.ID.EQ(UUID(sel))).
		LIMIT(1).
		QueryContext(ctx, r.conn, &modCat); err != nil {
		return nil, r.newErr(err, derr.NotFound)
	}

	return newDomainFromModel(&modCat), nil
}

func (r *catPostgresRepo) ListCats(
	ctx *context.Context,
) ([]domain.Cat, derr.DataError) {
	var modCatList []model.Cats
	if err := Cats.SELECT(Cats.AllColumns).
		QueryContext(ctx, r.conn, &modCatList); err != nil {
		return nil, r.newErr(err)
	}

	return util.MapRef(modCatList, newDomainFromModel), nil
}

func (r *catPostgresRepo) DeleteCat(
	ctx *context.Context,
	sel uuid.UUID,
) derr.DataError {
	if _, err := r.GetCat(ctx, sel); err != nil {
		return err
	}

	_, err := Cats.DELETE().
		WHERE(Cats.ID.EQ(UUID(sel))).
		ExecContext(ctx, r.conn)
	if err != nil {
		return r.newErr(err, derr.Conflict)
	}

	return nil
}

func (r *catPostgresRepo) newErr(
	err error,
	expTypes ...derr.DataErrorType,
) derr.DataError {
	errType, constr := commJet.ErrSpec(err)
	return derr.NewErr(errType, err, "cat", constr, expTypes...)
}
