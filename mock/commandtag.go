package mock

import (
	"fmt"
	"strings"
)

type CommandTag struct {
	Command      string
	AffectedRows int64
}

func (ct CommandTag) String() string      { return fmt.Sprintf("%s %d", ct.Command, ct.AffectedRows) }
func (ct CommandTag) RowsAffected() int64 { return ct.AffectedRows }
func (ct CommandTag) Insert() bool        { return strings.EqualFold(ct.Command, "Insert") }
func (ct CommandTag) Update() bool        { return strings.EqualFold(ct.Command, "Update") }
func (ct CommandTag) Delete() bool        { return strings.EqualFold(ct.Command, "Delete") }
func (ct CommandTag) Select() bool        { return strings.EqualFold(ct.Command, "Select") }
