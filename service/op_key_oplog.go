// Copyright 2018 The go-pttai Authors
// This file is part of the go-pttai library.
//
// The go-pttai library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-pttai library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-pttai library. If not, see <http://www.gnu.org/licenses/>.

package service

import (
	"github.com/ailabstw/go-pttai/common/types"
	"github.com/ailabstw/go-pttai/pttdb"
)

type OpKeyOplog struct {
	*BaseOplog `json:"O"`
}

func (o *OpKeyOplog) GetBaseOplog() *BaseOplog {
	return o.BaseOplog
}

func NewOpKeyOplog(keyID *types.PttID, ts types.Timestamp, doerID *types.PttID, op OpType, opData OpData, db *pttdb.LDBBatch, entityID *types.PttID, dbLock *types.LockMap) (*OpKeyOplog, error) {

	oplog, err := NewOplog(keyID, ts, doerID, op, opData, db, entityID, DBOpKeyOplogPrefix, DBOpKeyIdxOplogPrefix, nil, dbLock)
	if err != nil {
		return nil, err
	}

	return &OpKeyOplog{
		BaseOplog: oplog,
	}, nil
}

func (pm *BaseProtocolManager) NewOpKeyOplog(keyID *types.PttID, op OpType, opData OpData) (Oplog, error) {

	myID := pm.Ptt().GetMyEntity().GetID()
	entityID := pm.Entity().GetID()

	ts, err := types.GetTimestamp()
	if err != nil {
		return nil, err
	}

	return NewOpKeyOplog(keyID, ts, myID, op, opData, pm.DBOpKeyInfo(), entityID, pm.DBOpKeyLock())
}

func (pm *BaseProtocolManager) SetOpKeyDB(oplog *BaseOplog) {
	entityID := pm.Entity().GetID()
	oplog.SetDB(pm.DBOpKeyInfo(), entityID, DBOpKeyOplogPrefix, DBOpKeyIdxOplogPrefix, nil, pm.dbOpKeyLock)
}

func OplogsToOpKeyOplogs(oplogs []*BaseOplog) []*OpKeyOplog {
	typedLogs := make([]*OpKeyOplog, len(oplogs))
	for i, oplog := range oplogs {
		typedLogs[i] = &OpKeyOplog{BaseOplog: oplog}
	}
	return typedLogs
}

func OpKeyOplogsToOplogs(typedLogs []*OpKeyOplog) []*BaseOplog {
	oplogs := make([]*BaseOplog, len(typedLogs))
	for i, oplog := range typedLogs {
		oplogs[i] = oplog.BaseOplog
	}
	return oplogs
}

func OplogToOpKeyOplog(oplog *BaseOplog) *OpKeyOplog {
	if oplog == nil {
		return nil
	}
	return &OpKeyOplog{BaseOplog: oplog}
}
