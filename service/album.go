package service

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type AlbumService interface {
	GetAlbumListService(ctx context.Context) ([]*model.Album, error)
	GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.Album, error)
	PostAlbumService(ctx context.Context, album *model.Album) error
	DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error
	GetExtendAlbumService(ctx context.Context, albumID model.AlbumID) (*model.ExtendAlbum, error)
	GetExtendAlbumListService(ctx context.Context) ([]*model.ExtendAlbum, error)
}

type albumService struct {
	albumRepository repository.AlbumRepository
	singerService *singerService
}

var _ AlbumService = (*albumService)(nil)

func NewAlbumService(albumRepository repository.AlbumRepository, singerService *singerService) *albumService {
	return &albumService{albumRepository: albumRepository, singerService: singerService}
}

func (s *albumService) GetAlbumListService(ctx context.Context) ([]*model.Album, error) {
	albums, err := s.albumRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (s *albumService) GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.Album, error) {
	album, err := s.albumRepository.Get(ctx, albumID)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (s *albumService) PostAlbumService(ctx context.Context, album *model.Album) error {
	if err := s.albumRepository.Add(ctx, album); err != nil {
		return err
	}
	return nil
}

func (s *albumService) DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error {
	if err := s.albumRepository.Delete(ctx, albumID); err != nil {
		return err
	}
	return nil
}

func (s *albumService) GetExtendAlbumService(ctx context.Context, albumID model.AlbumID) (*model.ExtendAlbum, error) {
	album, err := s.albumRepository.Get(ctx, albumID)
	if err != nil {
		return nil, err
	}

	singer, err := s.singerService.GetSingerService(ctx, model.SingerID(album.SingerID))
	if err != nil {
		return nil, err
	}

	extendAlbum := ConvertExtend(ctx, album, singer)

	return extendAlbum, nil
}

func (s *albumService) GetExtendAlbumListService(ctx context.Context) ([]*model.ExtendAlbum, error) {
	albums, err := s.albumRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	extendAlbums := make([]*model.ExtendAlbum, 0, len(albums))
	for _, album := range albums {
		singer, err := s.singerService.GetSingerService(ctx, model.SingerID(album.SingerID))
		if err != nil {
			return nil, err
		}
		extendAlbum := ConvertExtend(ctx, album, singer)
		extendAlbums = append(extendAlbums, extendAlbum)
	}

	return extendAlbums, nil
}

func ConvertExtend(ctx context.Context, album *model.Album, singer *model.Singer) *model.ExtendAlbum {
	extendAlbum := model.ExtendAlbum{
		ID: album.ID,
		Title: album.Title,
		Singer: singer,
	}
	return &extendAlbum
}
