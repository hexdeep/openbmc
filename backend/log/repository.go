package log

import (
	"context"
	"fmt"

	"github.com/hexdeep/openbmc/backend/utils"
	"gorm.io/gorm"
)

type MySQL struct {
	DB        gorm.Interface[Log]
	Paginator *utils.Paginator
}

func NewMySQL(db *gorm.DB, paginator *utils.Paginator) *MySQL {
	return &MySQL{
		DB:        gorm.G[Log](db),
		Paginator: paginator,
	}
}

func (m *MySQL) Filter(filter *LogFilter) gorm.ChainInterface[Log] {

	q := m.DB.Scopes()

	if !filter.From.IsZero() {
		q = q.Where("created_at > ?", filter.From)
	}

	if !filter.To.IsZero() {
		q = q.Where("created_at < ?", filter.To)
	}

	if filter.ClientIP != "" {
		q = q.Where("client_ip = ?", filter.ClientIP)
	}

	if filter.Method != "" {
		q = q.Where("method = ?", filter.Method)
	}

	if filter.Path != "" {
		q = q.Where("path = ?", filter.Path)
	}

	if filter.Status != 0 {
		q = q.Where("status = ?", filter.Status)
	}

	return q
}

func (m *MySQL) Find(ctx context.Context, filter *LogFilter) ([]Log, int64, error) {

	q := m.Filter(filter)

	total, err := q.Count(ctx, "*")
	if err != nil {
		return nil, 0, fmt.Errorf("faild to count logs: %w", err)
	}

	data, err := q.Scopes(m.Paginator.Paginate(filter.Page, filter.Size)).Find(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get logs: %w", err)
	}

	return data, total, nil
}

func (m *MySQL) Create(ctx context.Context, log *Log) error {
	return m.DB.Create(ctx, log)
}

func (m *MySQL) Delete(ctx context.Context, filter *LogFilter) (int, error) {

	row, err := m.Filter(filter).Delete(ctx)
	if err != nil {
		return 0, err
	}

	return row, nil
}
