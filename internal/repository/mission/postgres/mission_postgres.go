package postgres

import (
	"database/sql"
	"sca/.gen/jet/sca/public/model"
	. "sca/.gen/jet/sca/public/table"
	"sca/internal/context"
	domain "sca/internal/domain/mission"
	derr "sca/internal/error/data"
	commJet "sca/internal/repository/common/postgres_jet"
	abstract "sca/internal/repository/mission"
	"sca/internal/util"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

type missionPostgresRepo struct {
	conn *sql.DB
}

func NewMissionPostgresRepo(conn *sql.DB) abstract.MissionRepo {
	return &missionPostgresRepo{
		conn,
	}
}

func (r *missionPostgresRepo) CreateMission(
	ctx *context.Context,
	create *domain.CreateMission,
) (*domain.Mission, derr.DataError) {
	tx, err := r.conn.Begin()
	if err != nil {
		return nil, r.newErr(err)
	}

	modMis := newModelFromCreate(create)
	if err := Missions.INSERT(Missions.Status).
		MODEL(modMis).
		RETURNING(Missions.AllColumns).QueryContext(ctx, tx, &modMis); err != nil {
		tx.Rollback()
		ctx.Error(err.Error())
		return nil, r.newErr(err)
	}

	insertTargList := newTargListFromCreate(modMis.ID, create.Targets)
	returnTargList := make([]model.Targets, 0, len(insertTargList))
	if err := Targets.INSERT(Targets.Name, Targets.Country, Targets.Notes, Targets.Status, Targets.MissionID).
		MODELS(insertTargList).
		RETURNING(Targets.AllColumns).QueryContext(ctx, tx, &returnTargList); err != nil {
		tx.Rollback()
		ctx.Error(err.Error())
		return nil, r.newErr(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, r.newErr(err)
	}

	return newDomainFromModel(&SaturatedMission{Missions: modMis, Targets: returnTargList}), nil
}

func (r *missionPostgresRepo) UpdateMission(
	ctx *context.Context,
	update *domain.UpdateMission,
) (*domain.Mission, derr.DataError) {
	modMis, err := newModelFromUpdate(update)
	if err != nil {
		return nil, r.newErr(err)
	}
	cols := make(ColumnList, 0)
	if update.Status != nil {
		cols = append(cols, Missions.Status)
	}
	if update.CatId != nil {
		cols = append(cols, Missions.CatID)
	}

	if err := Missions.UPDATE(cols).
		MODEL(modMis).
		WHERE(Missions.ID.EQ(UUID(update.Id))).
		RETURNING(Missions.AllColumns).QueryContext(ctx, r.conn, &modMis); err != nil {
		return nil, r.newErr(err, derr.Conflict, derr.NotFound)
	}

	return r.GetMission(ctx, update.Id)
}

func (r *missionPostgresRepo) GetMission(
	ctx *context.Context,
	sel uuid.UUID,
) (*domain.Mission, derr.DataError) {
	var dest SaturatedMission
	if err := Missions.SELECT(Missions.AllColumns, Targets.AllColumns).
		FROM(Targets.INNER_JOIN(Missions, Targets.MissionID.EQ(Missions.ID))).
		WHERE(Missions.ID.EQ(UUID(sel))).
		LIMIT(1).
		QueryContext(ctx, r.conn, &dest); err != nil {
		return nil, r.newErr(err, derr.NotFound)
	}

	return newDomainFromModel(&dest), nil
}

func (r *missionPostgresRepo) ListMissions(
	ctx *context.Context,
) ([]domain.Mission, derr.DataError) {
	var dest []SaturatedMission
	if err := Missions.SELECT(Missions.AllColumns, Targets.AllColumns).
		FROM(Targets.INNER_JOIN(Missions, Targets.MissionID.EQ(Missions.ID))).QueryContext(ctx, r.conn, &dest); err != nil {
		return nil, r.newErr(err)
	}

	return util.MapRef(dest, newDomainFromModel), nil
}

func (r *missionPostgresRepo) DeleteMission(
	ctx *context.Context,
	sel uuid.UUID,
) derr.DataError {
	if _, err := r.GetMission(ctx, sel); err != nil {
		return err
	}

	_, err := Missions.DELETE().
		WHERE(Missions.ID.EQ(UUID(sel))).
		Exec(r.conn)
	if err != nil {
		return r.newErr(err, derr.Conflict)
	}

	return nil
}

func (r *missionPostgresRepo) newErr(
	err error,
	expTypes ...derr.DataErrorType,
) derr.DataError {
	errType, constr := commJet.ErrSpec(err)
	return derr.NewErr(errType, err, "mission", constr, expTypes...)
}
