package repository

import "acne-scan-api/internal/model/domain"

func (articleRepository *ArticleRepositoryImpl) Create(article *domain.Article) error {
	stmt, err := articleRepository.DB.Prepare("insert into article (article_id,image,description) values (?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(article.ArticleId, article.Image, article.Description)
	if err != nil {
		return err
	}

	return nil
}
