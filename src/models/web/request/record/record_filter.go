package recordRequest

type RecordFilter struct {
	IdentityNumber *int64 `query:"identityDetail.identityNumber"`
	UserId         string `query:"createdBy.userId"`
	Nip            string `query:"createdBy.nip"`
	Limit          int    `query:"limit"`
	Offset         int    `query:"offset"`
	CreatedAt      string `query:"createdAt"`
}
