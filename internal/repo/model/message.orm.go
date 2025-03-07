package model

// !!! DO NOT EDIT THIS FILE

import (
	"context"
	"encoding/json"
	"github.com/iancoleman/strcase"
	"github.com/mylxsw/eloquent/query"
	"gopkg.in/guregu/null.v3"
	"time"
)

func init() {

}

// ChatMessagesN is a ChatMessages object, all fields are nullable
type ChatMessagesN struct {
	original          *chatMessagesOriginal
	chatMessagesModel *ChatMessagesModel

	Id            null.Int    `json:"id"`
	UserId        null.Int    `json:"user_id,omitempty"`
	RoomId        null.Int    `json:"room_id,omitempty"`
	Message       null.String `json:"message,omitempty"`
	Role          null.Int    `json:"role,omitempty"`
	TokenConsumed null.Int    `json:"token_consumed,omitempty"`
	QuotaConsumed null.Int    `json:"quota_consumed,omitempty"`
	Pid           null.Int    `json:"pid,omitempty"`
	Model         null.String `json:"model,omitempty"`
	CreatedAt     null.Time
	UpdatedAt     null.Time
}

// As convert object to other type
// dst must be a pointer to struct
func (inst *ChatMessagesN) As(dst interface{}) error {
	return query.Copy(inst, dst)
}

// SetModel set model for ChatMessages
func (inst *ChatMessagesN) SetModel(chatMessagesModel *ChatMessagesModel) {
	inst.chatMessagesModel = chatMessagesModel
}

// chatMessagesOriginal is an object which stores original ChatMessages from database
type chatMessagesOriginal struct {
	Id            null.Int
	UserId        null.Int
	RoomId        null.Int
	Message       null.String
	Role          null.Int
	TokenConsumed null.Int
	QuotaConsumed null.Int
	Pid           null.Int
	Model         null.String
	CreatedAt     null.Time
	UpdatedAt     null.Time
}

// Staled identify whether the object has been modified
func (inst *ChatMessagesN) Staled(onlyFields ...string) bool {
	if inst.original == nil {
		inst.original = &chatMessagesOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			return true
		}
		if inst.UserId != inst.original.UserId {
			return true
		}
		if inst.RoomId != inst.original.RoomId {
			return true
		}
		if inst.Message != inst.original.Message {
			return true
		}
		if inst.Role != inst.original.Role {
			return true
		}
		if inst.TokenConsumed != inst.original.TokenConsumed {
			return true
		}
		if inst.QuotaConsumed != inst.original.QuotaConsumed {
			return true
		}
		if inst.Pid != inst.original.Pid {
			return true
		}
		if inst.Model != inst.original.Model {
			return true
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			return true
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			return true
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					return true
				}
			case "user_id":
				if inst.UserId != inst.original.UserId {
					return true
				}
			case "room_id":
				if inst.RoomId != inst.original.RoomId {
					return true
				}
			case "message":
				if inst.Message != inst.original.Message {
					return true
				}
			case "role":
				if inst.Role != inst.original.Role {
					return true
				}
			case "token_consumed":
				if inst.TokenConsumed != inst.original.TokenConsumed {
					return true
				}
			case "quota_consumed":
				if inst.QuotaConsumed != inst.original.QuotaConsumed {
					return true
				}
			case "pid":
				if inst.Pid != inst.original.Pid {
					return true
				}
			case "model":
				if inst.Model != inst.original.Model {
					return true
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					return true
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					return true
				}
			default:
			}
		}
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *ChatMessagesN) StaledKV(onlyFields ...string) query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &chatMessagesOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			kv["id"] = inst.Id
		}
		if inst.UserId != inst.original.UserId {
			kv["user_id"] = inst.UserId
		}
		if inst.RoomId != inst.original.RoomId {
			kv["room_id"] = inst.RoomId
		}
		if inst.Message != inst.original.Message {
			kv["message"] = inst.Message
		}
		if inst.Role != inst.original.Role {
			kv["role"] = inst.Role
		}
		if inst.TokenConsumed != inst.original.TokenConsumed {
			kv["token_consumed"] = inst.TokenConsumed
		}
		if inst.QuotaConsumed != inst.original.QuotaConsumed {
			kv["quota_consumed"] = inst.QuotaConsumed
		}
		if inst.Pid != inst.original.Pid {
			kv["pid"] = inst.Pid
		}
		if inst.Model != inst.original.Model {
			kv["model"] = inst.Model
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			kv["created_at"] = inst.CreatedAt
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			kv["updated_at"] = inst.UpdatedAt
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					kv["id"] = inst.Id
				}
			case "user_id":
				if inst.UserId != inst.original.UserId {
					kv["user_id"] = inst.UserId
				}
			case "room_id":
				if inst.RoomId != inst.original.RoomId {
					kv["room_id"] = inst.RoomId
				}
			case "message":
				if inst.Message != inst.original.Message {
					kv["message"] = inst.Message
				}
			case "role":
				if inst.Role != inst.original.Role {
					kv["role"] = inst.Role
				}
			case "token_consumed":
				if inst.TokenConsumed != inst.original.TokenConsumed {
					kv["token_consumed"] = inst.TokenConsumed
				}
			case "quota_consumed":
				if inst.QuotaConsumed != inst.original.QuotaConsumed {
					kv["quota_consumed"] = inst.QuotaConsumed
				}
			case "pid":
				if inst.Pid != inst.original.Pid {
					kv["pid"] = inst.Pid
				}
			case "model":
				if inst.Model != inst.original.Model {
					kv["model"] = inst.Model
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					kv["created_at"] = inst.CreatedAt
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					kv["updated_at"] = inst.UpdatedAt
				}
			default:
			}
		}
	}

	return kv
}

// Save create a new model or update it
func (inst *ChatMessagesN) Save(ctx context.Context, onlyFields ...string) error {
	if inst.chatMessagesModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.chatMessagesModel.SaveOrUpdate(ctx, *inst, onlyFields...)
	if err != nil {
		return err
	}

	inst.Id = null.IntFrom(id)
	return nil
}

// Delete remove a chat_messages
func (inst *ChatMessagesN) Delete(ctx context.Context) error {
	if inst.chatMessagesModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.chatMessagesModel.DeleteById(ctx, inst.Id.Int64)
	if err != nil {
		return err
	}

	return nil
}

// String convert instance to json string
func (inst *ChatMessagesN) String() string {
	rs, _ := json.Marshal(inst)
	return string(rs)
}

type chatMessagesScope struct {
	name  string
	apply func(builder query.Condition)
}

var chatMessagesGlobalScopes = make([]chatMessagesScope, 0)
var chatMessagesLocalScopes = make([]chatMessagesScope, 0)

// AddGlobalScopeForChatMessages assign a global scope to a model
func AddGlobalScopeForChatMessages(name string, apply func(builder query.Condition)) {
	chatMessagesGlobalScopes = append(chatMessagesGlobalScopes, chatMessagesScope{name: name, apply: apply})
}

// AddLocalScopeForChatMessages assign a local scope to a model
func AddLocalScopeForChatMessages(name string, apply func(builder query.Condition)) {
	chatMessagesLocalScopes = append(chatMessagesLocalScopes, chatMessagesScope{name: name, apply: apply})
}

func (m *ChatMessagesModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range chatMessagesGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range chatMessagesLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *ChatMessagesModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *ChatMessagesModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}

	return true
}

type ChatMessages struct {
	Id            int64  `json:"id"`
	UserId        int64  `json:"user_id,omitempty"`
	RoomId        int64  `json:"room_id,omitempty"`
	Message       string `json:"message,omitempty"`
	Role          int64  `json:"role,omitempty"`
	TokenConsumed int64  `json:"token_consumed,omitempty"`
	QuotaConsumed int64  `json:"quota_consumed,omitempty"`
	Pid           int64  `json:"pid,omitempty"`
	Model         string `json:"model,omitempty"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (w ChatMessages) ToChatMessagesN(allows ...string) ChatMessagesN {
	if len(allows) == 0 {
		return ChatMessagesN{

			Id:            null.IntFrom(int64(w.Id)),
			UserId:        null.IntFrom(int64(w.UserId)),
			RoomId:        null.IntFrom(int64(w.RoomId)),
			Message:       null.StringFrom(w.Message),
			Role:          null.IntFrom(int64(w.Role)),
			TokenConsumed: null.IntFrom(int64(w.TokenConsumed)),
			QuotaConsumed: null.IntFrom(int64(w.QuotaConsumed)),
			Pid:           null.IntFrom(int64(w.Pid)),
			Model:         null.StringFrom(w.Model),
			CreatedAt:     null.TimeFrom(w.CreatedAt),
			UpdatedAt:     null.TimeFrom(w.UpdatedAt),
		}
	}

	res := ChatMessagesN{}
	for _, al := range allows {
		switch strcase.ToSnake(al) {

		case "id":
			res.Id = null.IntFrom(int64(w.Id))
		case "user_id":
			res.UserId = null.IntFrom(int64(w.UserId))
		case "room_id":
			res.RoomId = null.IntFrom(int64(w.RoomId))
		case "message":
			res.Message = null.StringFrom(w.Message)
		case "role":
			res.Role = null.IntFrom(int64(w.Role))
		case "token_consumed":
			res.TokenConsumed = null.IntFrom(int64(w.TokenConsumed))
		case "quota_consumed":
			res.QuotaConsumed = null.IntFrom(int64(w.QuotaConsumed))
		case "pid":
			res.Pid = null.IntFrom(int64(w.Pid))
		case "model":
			res.Model = null.StringFrom(w.Model)
		case "created_at":
			res.CreatedAt = null.TimeFrom(w.CreatedAt)
		case "updated_at":
			res.UpdatedAt = null.TimeFrom(w.UpdatedAt)
		default:
		}
	}

	return res
}

// As convert object to other type
// dst must be a pointer to struct
func (w ChatMessages) As(dst interface{}) error {
	return query.Copy(w, dst)
}

func (w *ChatMessagesN) ToChatMessages() ChatMessages {
	return ChatMessages{

		Id:            w.Id.Int64,
		UserId:        w.UserId.Int64,
		RoomId:        w.RoomId.Int64,
		Message:       w.Message.String,
		Role:          w.Role.Int64,
		TokenConsumed: w.TokenConsumed.Int64,
		QuotaConsumed: w.QuotaConsumed.Int64,
		Pid:           w.Pid.Int64,
		Model:         w.Model.String,
		CreatedAt:     w.CreatedAt.Time,
		UpdatedAt:     w.UpdatedAt.Time,
	}
}

// ChatMessagesModel is a model which encapsulates the operations of the object
type ChatMessagesModel struct {
	db        *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes  []string

	query query.SQLBuilder
}

var chatMessagesTableName = "chat_messages"

// ChatMessagesTable return table name for ChatMessages
func ChatMessagesTable() string {
	return chatMessagesTableName
}

const (
	FieldChatMessagesId            = "id"
	FieldChatMessagesUserId        = "user_id"
	FieldChatMessagesRoomId        = "room_id"
	FieldChatMessagesMessage       = "message"
	FieldChatMessagesRole          = "role"
	FieldChatMessagesTokenConsumed = "token_consumed"
	FieldChatMessagesQuotaConsumed = "quota_consumed"
	FieldChatMessagesPid           = "pid"
	FieldChatMessagesModel         = "model"
	FieldChatMessagesCreatedAt     = "created_at"
	FieldChatMessagesUpdatedAt     = "updated_at"
)

// ChatMessagesFields return all fields in ChatMessages model
func ChatMessagesFields() []string {
	return []string{
		"id",
		"user_id",
		"room_id",
		"message",
		"role",
		"token_consumed",
		"quota_consumed",
		"pid",
		"model",
		"created_at",
		"updated_at",
	}
}

func SetChatMessagesTable(tableName string) {
	chatMessagesTableName = tableName
}

// NewChatMessagesModel create a ChatMessagesModel
func NewChatMessagesModel(db query.Database) *ChatMessagesModel {
	return &ChatMessagesModel{
		db:                  query.NewDatabaseWrap(db),
		tableName:           chatMessagesTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes:  make([]string, 0),
		query:               query.Builder(),
	}
}

// GetDB return database instance
func (m *ChatMessagesModel) GetDB() query.Database {
	return m.db.GetDB()
}

func (m *ChatMessagesModel) clone() *ChatMessagesModel {
	return &ChatMessagesModel{
		db:                  m.db,
		tableName:           m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes:  append([]string{}, m.includeLocalScopes...),
		query:               m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *ChatMessagesModel) WithoutGlobalScopes(names ...string) *ChatMessagesModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *ChatMessagesModel) WithLocalScopes(names ...string) *ChatMessagesModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Condition add query builder to model
func (m *ChatMessagesModel) Condition(builder query.SQLBuilder) *ChatMessagesModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *ChatMessagesModel) Find(ctx context.Context, id int64) (*ChatMessagesN, error) {
	return m.First(ctx, m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *ChatMessagesModel) Exists(ctx context.Context, builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(ctx, builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *ChatMessagesModel) Count(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {
	sqlStr, params := m.query.
		Merge(builders...).
		Table(m.tableName).
		AppendCondition(m.applyScope()).
		ResolveCount()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	rows.Next()
	var res int64
	if err := rows.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (m *ChatMessagesModel) Paginate(ctx context.Context, page int64, perPage int64, builders ...query.SQLBuilder) ([]ChatMessagesN, query.PaginateMeta, error) {
	if page <= 0 {
		page = 1
	}

	if perPage <= 0 {
		perPage = 15
	}

	meta := query.PaginateMeta{
		PerPage: perPage,
		Page:    page,
	}

	count, err := m.Count(ctx, builders...)
	if err != nil {
		return nil, meta, err
	}

	meta.Total = count
	meta.LastPage = count / perPage
	if count%perPage != 0 {
		meta.LastPage += 1
	}

	res, err := m.Get(ctx, append([]query.SQLBuilder{query.Builder().Limit(perPage).Offset((page - 1) * perPage)}, builders...)...)
	if err != nil {
		return res, meta, err
	}

	return res, meta, nil
}

// Get retrieve all results for given query
func (m *ChatMessagesModel) Get(ctx context.Context, builders ...query.SQLBuilder) ([]ChatMessagesN, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"id",
			"user_id",
			"room_id",
			"message",
			"role",
			"token_consumed",
			"quota_consumed",
			"pid",
			"model",
			"created_at",
			"updated_at",
		)
	}

	fields := b.GetFields()
	selectFields := make([]query.Expr, 0)

	for _, f := range fields {
		switch strcase.ToSnake(f.Value) {

		case "id":
			selectFields = append(selectFields, f)
		case "user_id":
			selectFields = append(selectFields, f)
		case "room_id":
			selectFields = append(selectFields, f)
		case "message":
			selectFields = append(selectFields, f)
		case "role":
			selectFields = append(selectFields, f)
		case "token_consumed":
			selectFields = append(selectFields, f)
		case "quota_consumed":
			selectFields = append(selectFields, f)
		case "pid":
			selectFields = append(selectFields, f)
		case "model":
			selectFields = append(selectFields, f)
		case "created_at":
			selectFields = append(selectFields, f)
		case "updated_at":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*ChatMessagesN, []interface{}) {
		var chatMessagesVar ChatMessagesN
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {

			case "id":
				scanFields = append(scanFields, &chatMessagesVar.Id)
			case "user_id":
				scanFields = append(scanFields, &chatMessagesVar.UserId)
			case "room_id":
				scanFields = append(scanFields, &chatMessagesVar.RoomId)
			case "message":
				scanFields = append(scanFields, &chatMessagesVar.Message)
			case "role":
				scanFields = append(scanFields, &chatMessagesVar.Role)
			case "token_consumed":
				scanFields = append(scanFields, &chatMessagesVar.TokenConsumed)
			case "quota_consumed":
				scanFields = append(scanFields, &chatMessagesVar.QuotaConsumed)
			case "pid":
				scanFields = append(scanFields, &chatMessagesVar.Pid)
			case "model":
				scanFields = append(scanFields, &chatMessagesVar.Model)
			case "created_at":
				scanFields = append(scanFields, &chatMessagesVar.CreatedAt)
			case "updated_at":
				scanFields = append(scanFields, &chatMessagesVar.UpdatedAt)
			}
		}

		return &chatMessagesVar, scanFields
	}

	sqlStr, params := b.Fields(selectFields...).ResolveQuery()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	chatMessagess := make([]ChatMessagesN, 0)
	for rows.Next() {
		chatMessagesReal, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		chatMessagesReal.original = &chatMessagesOriginal{}
		_ = query.Copy(chatMessagesReal, chatMessagesReal.original)

		chatMessagesReal.SetModel(m)
		chatMessagess = append(chatMessagess, *chatMessagesReal)
	}

	return chatMessagess, nil
}

// First return first result for given query
func (m *ChatMessagesModel) First(ctx context.Context, builders ...query.SQLBuilder) (*ChatMessagesN, error) {
	res, err := m.Get(ctx, append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, query.ErrNoResult
	}

	return &res[0], nil
}

// Create save a new chat_messages to database
func (m *ChatMessagesModel) Create(ctx context.Context, kv query.KV) (int64, error) {

	if _, ok := kv["created_at"]; !ok {
		kv["created_at"] = time.Now()
	}

	if _, ok := kv["updated_at"]; !ok {
		kv["updated_at"] = time.Now()
	}

	sqlStr, params := m.query.Table(m.tableName).ResolveInsert(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// SaveAll save all chat_messagess to database
func (m *ChatMessagesModel) SaveAll(ctx context.Context, chatMessagess []ChatMessagesN) ([]int64, error) {
	ids := make([]int64, 0)
	for _, chatMessages := range chatMessagess {
		id, err := m.Save(ctx, chatMessages)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a chat_messages to database
func (m *ChatMessagesModel) Save(ctx context.Context, chatMessages ChatMessagesN, onlyFields ...string) (int64, error) {
	return m.Create(ctx, chatMessages.StaledKV(onlyFields...))
}

// SaveOrUpdate save a new chat_messages or update it when it has a id > 0
func (m *ChatMessagesModel) SaveOrUpdate(ctx context.Context, chatMessages ChatMessagesN, onlyFields ...string) (id int64, updated bool, err error) {
	if chatMessages.Id.Int64 > 0 {
		_, _err := m.UpdateById(ctx, chatMessages.Id.Int64, chatMessages, onlyFields...)
		return chatMessages.Id.Int64, true, _err
	}

	_id, _err := m.Save(ctx, chatMessages, onlyFields...)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *ChatMessagesModel) UpdateFields(ctx context.Context, kv query.KV, builders ...query.SQLBuilder) (int64, error) {
	if len(kv) == 0 {
		return 0, nil
	}

	kv["updated_at"] = time.Now()

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).
		Table(m.tableName).
		ResolveUpdate(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Update update a model for given query
func (m *ChatMessagesModel) Update(ctx context.Context, builder query.SQLBuilder, chatMessages ChatMessagesN, onlyFields ...string) (int64, error) {
	return m.UpdateFields(ctx, chatMessages.StaledKV(onlyFields...), builder)
}

// UpdateById update a model by id
func (m *ChatMessagesModel) UpdateById(ctx context.Context, id int64, chatMessages ChatMessagesN, onlyFields ...string) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).UpdateFields(ctx, chatMessages.StaledKV(onlyFields...))
}

// Delete remove a model
func (m *ChatMessagesModel) Delete(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).Table(m.tableName).ResolveDelete()

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

// DeleteById remove a model by id
func (m *ChatMessagesModel) DeleteById(ctx context.Context, id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Delete(ctx)
}
