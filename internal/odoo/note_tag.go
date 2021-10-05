package odoo

import (
	"fmt"
)

// NoteTag represents note.tag model.
type NoteTag struct {
	Color       *Int      `xmlrpc:"color,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// NoteTags represents array of note.tag model.
type NoteTags []NoteTag

// NoteTagModel is the odoo model name.
const NoteTagModel = "note.tag"

// Many2One convert NoteTag to *Many2One.
func (nt *NoteTag) Many2One() *Many2One {
	return NewMany2One(nt.Id.Get(), "")
}

// CreateNoteTag creates a new note.tag model and returns its id.
func (c *Client) CreateNoteTag(nt *NoteTag) (int64, error) {
	return c.Create(NoteTagModel, nt)
}

// UpdateNoteTag updates an existing note.tag record.
func (c *Client) UpdateNoteTag(nt *NoteTag) error {
	return c.UpdateNoteTags([]int64{nt.Id.Get()}, nt)
}

// UpdateNoteTags updates existing note.tag records.
// All records (represented by ids) will be updated by nt values.
func (c *Client) UpdateNoteTags(ids []int64, nt *NoteTag) error {
	return c.Update(NoteTagModel, ids, nt)
}

// DeleteNoteTag deletes an existing note.tag record.
func (c *Client) DeleteNoteTag(id int64) error {
	return c.DeleteNoteTags([]int64{id})
}

// DeleteNoteTags deletes existing note.tag records.
func (c *Client) DeleteNoteTags(ids []int64) error {
	return c.Delete(NoteTagModel, ids)
}

// GetNoteTag gets note.tag existing record.
func (c *Client) GetNoteTag(id int64) (*NoteTag, error) {
	nts, err := c.GetNoteTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if nts != nil && len(*nts) > 0 {
		return &((*nts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of note.tag not found", id)
}

// GetNoteTags gets note.tag existing records.
func (c *Client) GetNoteTags(ids []int64) (*NoteTags, error) {
	nts := &NoteTags{}
	if err := c.Read(NoteTagModel, ids, nil, nts); err != nil {
		return nil, err
	}
	return nts, nil
}

// FindNoteTag finds note.tag record by querying it with criteria.
func (c *Client) FindNoteTag(criteria *Criteria) (*NoteTag, error) {
	nts := &NoteTags{}
	if err := c.SearchRead(NoteTagModel, criteria, NewOptions().Limit(1), nts); err != nil {
		return nil, err
	}
	if nts != nil && len(*nts) > 0 {
		return &((*nts)[0]), nil
	}
	return nil, fmt.Errorf("note.tag was not found")
}

// FindNoteTags finds note.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindNoteTags(criteria *Criteria, options *Options) (*NoteTags, error) {
	nts := &NoteTags{}
	if err := c.SearchRead(NoteTagModel, criteria, options, nts); err != nil {
		return nil, err
	}
	return nts, nil
}

// FindNoteTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindNoteTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(NoteTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindNoteTagId finds record id by querying it with criteria.
func (c *Client) FindNoteTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(NoteTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("note.tag was not found")
}
