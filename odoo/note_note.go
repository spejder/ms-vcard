package odoo

import (
	"fmt"
)

// NoteNote represents note.note model.
type NoteNote struct {
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	Color                       *Int       `xmlrpc:"color,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateDone                    *Time      `xmlrpc:"date_done,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	Memo                        *String    `xmlrpc:"memo,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	Open                        *Bool      `xmlrpc:"open,omptempty"`
	Sequence                    *Int       `xmlrpc:"sequence,omptempty"`
	StageId                     *Many2One  `xmlrpc:"stage_id,omptempty"`
	StageIds                    *Relation  `xmlrpc:"stage_ids,omptempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omptempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// NoteNotes represents array of note.note model.
type NoteNotes []NoteNote

// NoteNoteModel is the odoo model name.
const NoteNoteModel = "note.note"

// Many2One convert NoteNote to *Many2One.
func (nn *NoteNote) Many2One() *Many2One {
	return NewMany2One(nn.Id.Get(), "")
}

// CreateNoteNote creates a new note.note model and returns its id.
func (c *Client) CreateNoteNote(nn *NoteNote) (int64, error) {
	return c.Create(NoteNoteModel, nn)
}

// UpdateNoteNote updates an existing note.note record.
func (c *Client) UpdateNoteNote(nn *NoteNote) error {
	return c.UpdateNoteNotes([]int64{nn.Id.Get()}, nn)
}

// UpdateNoteNotes updates existing note.note records.
// All records (represented by ids) will be updated by nn values.
func (c *Client) UpdateNoteNotes(ids []int64, nn *NoteNote) error {
	return c.Update(NoteNoteModel, ids, nn)
}

// DeleteNoteNote deletes an existing note.note record.
func (c *Client) DeleteNoteNote(id int64) error {
	return c.DeleteNoteNotes([]int64{id})
}

// DeleteNoteNotes deletes existing note.note records.
func (c *Client) DeleteNoteNotes(ids []int64) error {
	return c.Delete(NoteNoteModel, ids)
}

// GetNoteNote gets note.note existing record.
func (c *Client) GetNoteNote(id int64) (*NoteNote, error) {
	nns, err := c.GetNoteNotes([]int64{id})
	if err != nil {
		return nil, err
	}
	if nns != nil && len(*nns) > 0 {
		return &((*nns)[0]), nil
	}
	return nil, fmt.Errorf("id %v of note.note not found", id)
}

// GetNoteNotes gets note.note existing records.
func (c *Client) GetNoteNotes(ids []int64) (*NoteNotes, error) {
	nns := &NoteNotes{}
	if err := c.Read(NoteNoteModel, ids, nil, nns); err != nil {
		return nil, err
	}
	return nns, nil
}

// FindNoteNote finds note.note record by querying it with criteria.
func (c *Client) FindNoteNote(criteria *Criteria) (*NoteNote, error) {
	nns := &NoteNotes{}
	if err := c.SearchRead(NoteNoteModel, criteria, NewOptions().Limit(1), nns); err != nil {
		return nil, err
	}
	if nns != nil && len(*nns) > 0 {
		return &((*nns)[0]), nil
	}
	return nil, fmt.Errorf("note.note was not found")
}

// FindNoteNotes finds note.note records by querying it
// and filtering it with criteria and options.
func (c *Client) FindNoteNotes(criteria *Criteria, options *Options) (*NoteNotes, error) {
	nns := &NoteNotes{}
	if err := c.SearchRead(NoteNoteModel, criteria, options, nns); err != nil {
		return nil, err
	}
	return nns, nil
}

// FindNoteNoteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindNoteNoteIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(NoteNoteModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindNoteNoteId finds record id by querying it with criteria.
func (c *Client) FindNoteNoteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(NoteNoteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("note.note was not found")
}
