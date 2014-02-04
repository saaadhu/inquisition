package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "inquisition/data"
    "errors"
)

func getConnection() (con *sql.DB) {
    con, err := sql.Open ("mysql", "inq_writer:inquisition_writer_$#1@/inquisition")
    
    if err != nil {
        panic (err)
    }
    
    return
}

func AuthenticateTaker (username string, password string) (taker *data.Taker, err error) {
    con := getConnection()
    taker = nil

    rows, err := con.Query ("select id, name, login, department, college, test_spec_id " +
                            "from takers where " +
                            "takers.login=? and takers.password=? ",
                            username, password)

    if err != nil {
        return }

    if rows.Next() == false {
        err = errors.New("No such username or invalid password")
        return
    }


    taker = &data.Taker{};
    rows.Scan(taker.Id, taker.Name, taker.Login, taker.Department, taker.College, taker.TestSpecId)

    return
}



/*
func GetAllFeedsForUser (userid string) []data.Feed {
    con := getConnection()
    defer con.Close()

    rows, err := con.Query ("select feeds.id, feeds.title, feeds.url, server_last_modified, " +
                                "(select count(id) from feeditems " +
                                 "where userfeeds.feed = feeditems.feed and userfeeds.userid=? and " +
                                    "(feeditems.published > (select published from feeditems where id = last_read_item) or userfeeds.last_read_item IS NULL)" +
                                    "group by userfeeds.feed order by feeditems.published desc)," +
                             "last_fetch from feeds, userfeeds where userfeeds.userid=? and feeds.id = userfeeds.feed" , userid, userid)
    
    if err != nil {
        panic (err)
    }
    
    feeds := []data.Feed {}
    for rows.Next() {
        var id, url,title, serverLastModified string
        var lastFetch time.Time
        var unreadItemsCount int
        rows.Scan (&id, &title, &url, &serverLastModified, &unreadItemsCount, &lastFetch)
        feeds = append (feeds, data.Feed { Id : id, Title: title, Url : url, LastModified : serverLastModified, LastFetch : lastFetch, UnreadItemsCount : unreadItemsCount })
    }
    
    return feeds
}
*/
