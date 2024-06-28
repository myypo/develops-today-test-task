package postgres

import (
	"database/sql"
	"sca/.gen/jet/sca/public/model"
	. "sca/.gen/jet/sca/public/table"
	"sca/internal/context"
	domain "sca/internal/domain/target"
	derr "sca/internal/error/data"
	commJet "sca/internal/repository/common/postgres_jet"
	abstract "sca/internal/repository/target"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

type targetPostgresRepo struct {
	conn *sql.DB
}

func NewTargetPostgresRepo(conn *sql.DB) abstract.TargetRepo {
	return &targetPostgresRepo{
		conn,
	}
}

func (r *targetPostgresRepo) AddTarget(
	ctx *context.Context,
	add *domain.AddTarget,
) (*domain.Target, derr.DataError) {
	modTarg, err := newModelFromAdd(add)
	if err != nil {
		return nil, r.newErr(err)
	}

	if err := Targets.INSERT(Targets.AllColumns).
		MODEL(modTarg).
		RETURNING(Targets.AllColumns).
		Query(r.conn, &modTarg); err != nil {
		return nil, r.newErr(err)
	}

	return NewDomainFromModel(&modTarg), nil
}

func (r *targetPostgresRepo) UpdateTarget(
	ctx *context.Context,
	update *domain.UpdateTarget,
) (*domain.Target, derr.DataError) {
	modTarg := newModelFromUpdate(update)
	if err := Targets.UPDATE(Targets.Name, Targets.Country, Targets.Notes, Targets.Status).
		MODEL(modTarg).
		WHERE(Targets.ID.EQ(String(update.Id))).
		RETURNING(Targets.AllColumns).Query(r.conn, &modTarg); err != nil {
		return nil, r.newErr(err)
	}

	return NewDomainFromModel(&modTarg), nil
}

func (r *targetPostgresRepo) GetTarget(
	ctx *context.Context,
	sel uuid.UUID,
) (*domain.Target, derr.DataError) {
	var modTarg model.Targets
	if err := Targets.SELECT(Targets.AllColumns).
		WHERE(Targets.ID.EQ(UUID(sel))).
		LIMIT(1).
		Query(r.conn, &modTarg); err != nil {
		return nil, r.newErr(err, derr.NotFound)
	}

	return NewDomainFromModel(&modTarg), nil
}

func (r *targetPostgresRepo) DeleteTarget(
	ctx *context.Context,
	sel uuid.UUID,
) derr.DataError {
	if _, err := r.GetTarget(ctx, sel); err != nil {
		return err
	}

	_, err := Targets.DELETE().
		WHERE(Targets.ID.EQ(UUID(sel))).
		Exec(r.conn)
	if err != nil {
		return r.newErr(err, derr.Conflict)
	}

	return nil
}

func (r *targetPostgresRepo) newErr(
	err error,
	expTypes ...derr.DataErrorType,
) derr.DataError {
	errType, constr := commJet.ErrSpec(err)
	return derr.NewErr(errType, err, "target", constr, expTypes...)
}
