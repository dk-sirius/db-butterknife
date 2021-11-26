db-decl

针对postgresql数据库表声明，将golang struct 定义数据库的表转化为对应的数据库生Schema,同时完相关文件等。

安装

    // (需要把你的$GOPATH/bin 纳入系统path下)
    go install 
    // 其他方式（略）

Shell 使用

    db-decl gen -name your_basebase_name -t your_desc_object -f your desc_file

//go:generate 使用

    //go:generate db-decl gen -n wxf -t Account
    
    // Account 账户
    //@def primary f_id
    //@def unique_index i_userID f_userID
    //@def index i_name f_name
    //@def unique_index i_userID_name f_userID f_name
    type Account struct {
    	tmp.TimestampDeep
    	TimestampAt
    	AccountID
    	Name     string `db:"f_name,size=50,default=''"`
    	Password string `db:"f_password"`
    	UserID   uint64 `db:"f_userID"`
    	Nickname string `db:"f_nick_name,size=90,default=''"`
    }
    
    type AccountID struct {
    	ID uint64 `db:"f_id,autoincrement"`
    }






