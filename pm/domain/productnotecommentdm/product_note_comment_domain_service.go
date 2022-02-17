package productnotecommentdm

type productNoteCommentDomainService struct {
	productNoteCommentRepository ProductNoteCommentRepository
}

func NewProductNoteCommentDomainService(productNoteCommentRepository ProductNoteCommentRepository) *productNoteCommentDomainService {
	return &productNoteCommentDomainService{
		productNoteCommentRepository: productNoteCommentRepository,
	}
}
