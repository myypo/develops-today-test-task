package request

type InternalUUIDSelector struct {
	ID string `uri:"id" binding:"required,uuid"`
}
