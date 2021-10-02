package odoo

import (
	"fmt"
)

// NoteStage represents note.stage model.
type NoteStage struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Fold        *Bool     `xmlrpc:"fold,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	UserId      *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// NoteStages represents array of note.stage model.
type NoteStages []NoteStage

// NoteStageModel is the odoo model name.
const NoteStageModel = "note.stage"

// Many2One convert NoteStage to *Many2One.
func (ns *NoteStage) Many2One() *Many2One {
	return NewMany2One(ns.Id.Get(), "")
}

// CreateNoteStage creates a new note.stage model and returns its id.
func (c *Client) CreateNoteStage(ns *NoteStage) (int64, error) {
	return c.Create(NoteStageModel, ns)
}

// UpdateNoteStage updates an existing note.stage record.
func (c *Client) UpdateNoteStage(ns *NoteStage) error {
	return c.UpdateNoteStages([]int64{ns.Id.Get()}, ns)
}

// UpdateNoteStages updates existing note.stage records.
// All records (represented by ids) will be updated by ns values.
func (c *Client) UpdateNoteStages(ids []int64, ns *NoteStage) error {
	return c.Update(NoteStageModel, ids, ns)
}

// DeleteNoteStage deletes an existing note.stage record.
func (c *Client) DeleteNoteStage(id int64) error {
	return c.DeleteNoteStages([]int64{id})
}

// DeleteNoteStages deletes existing note.stage records.
func (c *Client) DeleteNoteStages(ids []int64) error {
	return c.Delete(NoteStageModel, ids)
}

// GetNoteStage gets note.stage existing record.
func (c *Client) GetNoteStage(id int64) (*NoteStage, error) {
	nss, err := c.GetNoteStages([]int64{id})
	if err != nil {
		return nil, err
	}
	if nss != nil && len(*nss) > 0 {
		return &((*nss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of note.stage not found", id)
}

// GetNoteStages gets note.stage existing records.
func (c *Client) GetNoteStages(ids []int64) (*NoteStages, error) {
	nss := &NoteStages{}
	if err := c.Read(NoteStageModel, ids, nil, nss); err != nil {
		return nil, err
	}
	return nss, nil
}

// FindNoteStage finds note.stage record by querying it with criteria.
func (c *Client) FindNoteStage(criteria *Criteria) (*NoteStage, error) {
	nss := &NoteStages{}
	if err := c.SearchRead(NoteStageModel, criteria, NewOptions().Limit(1), nss); err != nil {
		return nil, err
	}
	if nss != nil && len(*nss) > 0 {
		return &((*nss)[0]), nil
	}
	return nil, fmt.Errorf("note.stage was not found")
}

// FindNoteStages finds note.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindNoteStages(criteria *Criteria, options *Options) (*NoteStages, error) {
	nss := &NoteStages{}
	if err := c.SearchRead(NoteStageModel, criteria, options, nss); err != nil {
		return nil, err
	}
	return nss, nil
}

// FindNoteStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindNoteStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(NoteStageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindNoteStageId finds record id by querying it with criteria.
func (c *Client) FindNoteStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(NoteStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("note.stage was not found")
}
