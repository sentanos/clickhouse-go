package clickhouse

import (
	"time"

	"github.com/kshvakov/clickhouse/lib/data"
	"github.com/kshvakov/clickhouse/lib/protocol"
)

func (ch *clickhouse) readBlock() (*data.Block, error) {
	{
		ch.conn.SetReadDeadline(time.Now().Add(ch.readTimeout))
		ch.conn.SetWriteDeadline(time.Now().Add(ch.writeTimeout))
	}
	if ch.ServerInfo.Revision >= protocol.DBMS_MIN_REVISION_WITH_TEMPORARY_TABLES {
		if _, err := ch.decoder.String(); err != nil {
			return nil, err
		}
	}
	if ch.compress {

	}
	var block data.Block
	if err := block.Read(&ch.ServerInfo, ch.decoder); err != nil {
		return nil, err
	}
	return &block, nil
}
