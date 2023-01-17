package url

import (
	"errors"
	"gorm.io/gorm"
	"shortened_link/model"
	"time"
)

type PostgresUrlServiceImpl struct {
	DB *gorm.DB
}

func (p *PostgresUrlServiceImpl) AddUrl(srcUrl string) (*model.ShortedUrl, error) {
	var shortUrl model.ShortedUrl
	newShort := p.generateShortedUrl()
	shortUrl.ShortedUrl = newShort

	expireTime := time.Now().Add(24 * 7 * time.Hour)
	shortUrl = model.ShortedUrl{
		LongUrl:    srcUrl,
		ShortedUrl: newShort,
		ExpiredAt:  expireTime,
		Custom:     false,
	}
	err := p.DB.Create(&shortUrl).Error
	return &shortUrl, err
}

func (p *PostgresUrlServiceImpl) AddCustomUrl(customUrl, srcUrl string) (*model.ShortedUrl, error) {
	var shortUrl model.ShortedUrl
	var count int64
	shortUrl.ShortedUrl = customUrl
	p.DB.Where("shorted_url =?", customUrl).Count(&count)
	if count == 0 {
		expireTime := time.Now().Add(24 * 7 * time.Hour)
		shortUrl = model.ShortedUrl{
			LongUrl:    srcUrl,
			ShortedUrl: customUrl,
			ExpiredAt:  expireTime,
			Custom:     true,
		}

	} else {
		return &shortUrl, errors.New("this short url is already exist! try another")
	}
	err := p.DB.Create(&shortUrl).Error
	return &shortUrl, err
}

func (p *PostgresUrlServiceImpl) GetUrl(shortUrl string) (*model.ShortedUrl, bool) {
	var shortedUrl model.ShortedUrl
	var isFound = true
	err := p.DB.Where("shorted_url =?", shortUrl).First(&shortedUrl).Error
	if err != nil {
		isFound = false
		return nil, isFound
	}
	if isFound && shortedUrl.IsExpire() {
		p.DB.Delete(&shortedUrl)
		return &shortedUrl, !isFound
	}

	return &shortedUrl, isFound
}

func (p *PostgresUrlServiceImpl) UpdateLongUrl(shorted, newLong string) (*model.ShortedUrl, error) {
	var shortedUrl model.ShortedUrl
	p.DB.Where("shorted_url=?", shorted).Find(&shortedUrl)
	shortedUrl.LongUrl = newLong
	err := p.DB.Where("shorted_url = ?", shorted).Updates(&shortedUrl).Error
	if err != nil {
		return nil, err
	}
	return &shortedUrl, err
}

func (p *PostgresUrlServiceImpl) UpdateShortUrl(key, newShort string) (*model.ShortedUrl, error) {
	var shortedUrl model.ShortedUrl
	p.DB.Where("shorted_url =?", key).Find(&shortedUrl)
	ex := shortedUrl.ExpiredAt
	y := shortedUrl.LongUrl
	var count int64
	//check new short is unique
	p.DB.Where("shorted_url =?", newShort).Count(&count)
	if count > 0 {
		return nil, errors.New("this short Url is already exist")
	}
	err := p.DB.Where("shorted_url=?", key).Delete(&shortedUrl).Error
	if err != nil {
		return nil, err
	}

	shortedUrl = model.ShortedUrl{
		LongUrl:    y,
		ShortedUrl: newShort,
		ExpiredAt:  ex,
		Custom:     true,
	}
	p.DB.Where("shorted_url = ?", key).Save(&shortedUrl)

	return &shortedUrl, err
}

func (p *PostgresUrlServiceImpl) DeleteUrl(shorted string) error {
	var shortedUrl model.ShortedUrl
	return p.DB.Where("shorted_url=?", shorted).Delete(&shortedUrl).Error
}

func (p *PostgresUrlServiceImpl) generateShortedUrl() string {
	lenght := int64(len(alphabet))
	var count int64
	shortUrl := ""
	p.DB.Model(&model.ShortedUrl{}).Count(&count)
	for count > 0 {
		i := count % lenght
		shortUrl += string(alphabet[i])
		count = (count - i) / lenght
	}
	return shortUrl
}
