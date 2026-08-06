package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"requestEth/internal/biz"
	"strings"
	"time"
)

type SwapTrade struct {
	ID uint64 `gorm:"primarykey;type:bigint unsigned;comment:主键"`

	BlockNumber uint64 `gorm:"type:bigint unsigned;not null;index:idx_block_number;index:idx_block_log,priority:1;comment:区块高度"`
	LogIndex    uint32 `gorm:"type:int unsigned;not null;default:0;index:idx_block_log,priority:2;comment:logIndex(排序用)"`
	BlockTime   uint64 `gorm:"type:bigint unsigned;not null;index:idx_block_number;index:idx_block_log,priority:1;comment:区块高度"`

	Sender string `gorm:"type:varchar(42);not null;index:idx_sender_block,priority:1;comment:事件sender(0x...)"`
	ToAddr string `gorm:"column:to_addr;type:varchar(42);not null;index:idx_to_block,priority:1;comment:事件to(0x...)"`

	Side uint8 `gorm:"type:tinyint unsigned;not null;default:0;index:idx_side_block,priority:1;comment:方向1=BUY(DL)2=SELL(DL)0=UNKNOWN"`

	Amount0In       float64 `gorm:"column:amount0_in;type:decimal(65,18);not null;default:0;comment:DL in"`
	Amount1In       float64 `gorm:"column:amount1_in;type:decimal(65,18);not null;default:0;comment:OTHER in"`
	Amount0OutGross float64 `gorm:"column:amount0_out_gross;type:decimal(65,18);not null;default:0;comment:DL out"`
	Amount0OutNet   float64 `gorm:"column:amount0_out_net;type:decimal(65,18);not null;default:0;comment:DL out net"`
	Amount1OutGross float64 `gorm:"column:amount1_out_gross;type:decimal(65,18);not null;default:0;comment:OTHER out"`

	CreatedAt time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;comment:更新时间"`
}

type PrimaryBuy struct {
	ID uint64 `gorm:"primarykey;type:bigint unsigned;comment:主键"`

	BlockNumber uint64 `gorm:"column:block_number;type:bigint unsigned;not null;index:idx_block_log,priority:1;comment:区块高度"`
	BlockTime   uint64 `gorm:"column:block_time;type:bigint unsigned;not null;index:idx_block_time;comment:区块时间(秒)"`
	LogIndex    uint32 `gorm:"column:log_index;type:int unsigned;not null;default:0;index:idx_block_log,priority:2;comment:logIndex(排序用)"`

	Buyer  string `gorm:"column:buyer;type:varchar(42);not null;index:idx_buyer;comment:buyer"`
	ToAddr string `gorm:"column:to_addr;type:varchar(42);not null;index:idx_to;comment:to"`

	UsdtUsed     float64 `gorm:"column:usdt_used;type:decimal(65,18);not null;default:0;comment:usdtUsed"`
	AusdGrossOut float64 `gorm:"column:ausd_gross_out;type:decimal(65,18);not null;default:0;comment:ausdGrossOut"`
	AusdFee      float64 `gorm:"column:ausd_fee;type:decimal(65,18);not null;default:0;comment:ausdFee"`
	AusdNetOut   float64 `gorm:"column:ausd_net_out;type:decimal(65,18);not null;default:0;comment:ausdNetOut"`
	PriceBefore  float64 `gorm:"column:price_before;type:decimal(65,18);not null;default:0;comment:priceBefore"`
	PriceAfter   float64 `gorm:"column:price_after;type:decimal(65,18);not null;default:0;comment:priceAfter"`

	CreatedAt time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;comment:更新时间"`
}

type PrimarySell struct {
	ID uint64 `gorm:"primarykey;type:bigint unsigned;comment:主键"`

	BlockNumber uint64 `gorm:"column:block_number;type:bigint unsigned;not null;index:idx_block_log,priority:1;comment:区块高度"`
	BlockTime   uint64 `gorm:"column:block_time;type:bigint unsigned;not null;index:idx_block_time;comment:区块时间(秒)"`
	LogIndex    uint32 `gorm:"column:log_index;type:int unsigned;not null;default:0;index:idx_block_log,priority:2;comment:logIndex(排序用)"`

	Seller string `gorm:"column:seller;type:varchar(42);not null;index:idx_seller;comment:seller"`
	ToAddr string `gorm:"column:to_addr;type:varchar(42);not null;index:idx_to;comment:to"`

	AusdGrossIn float64 `gorm:"column:ausd_gross_in;type:decimal(65,18);not null;default:0;comment:ausdGrossIn"`
	AusdFee     float64 `gorm:"column:ausd_fee;type:decimal(65,18);not null;default:0;comment:ausdFee"`
	AusdBurn    float64 `gorm:"column:ausd_burn;type:decimal(65,18);not null;default:0;comment:ausdBurn"`
	UsdtOut     float64 `gorm:"column:usdt_out;type:decimal(65,18);not null;default:0;comment:usdtOut"`
	PriceBefore float64 `gorm:"column:price_before;type:decimal(65,18);not null;default:0;comment:priceBefore"`
	PriceAfter  float64 `gorm:"column:price_after;type:decimal(65,18);not null;default:0;comment:priceAfter"`

	CreatedAt time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;comment:更新时间"`
}

type RewardNotified struct {
	ID uint64 `gorm:"primarykey;type:bigint unsigned;comment:主键"`

	BlockNumber uint64 `gorm:"column:block_number;type:bigint unsigned;not null;index:idx_block_log,priority:1;comment:区块高度"`
	BlockTime   uint64 `gorm:"column:block_time;type:bigint unsigned;not null;index:idx_block_time;comment:区块时间(秒)"`
	LogIndex    uint32 `gorm:"column:log_index;type:int unsigned;not null;default:0;index:idx_block_log,priority:2;comment:logIndex(排序用)"`

	User string `gorm:"column:user;type:varchar(42);not null;index:idx_user_time,priority:1;comment:indexed user"`
	L1   string `gorm:"column:l1;type:varchar(42);not null;index:idx_l1_time,priority:1;comment:indexed l1"`
	L2   string `gorm:"column:l2;type:varchar(42);not null;index:idx_l2_time,priority:1;comment:indexed l2"`

	Profit    float64 `gorm:"column:profit;type:decimal(65,18);not null;default:0;comment:profit"`
	UserShare float64 `gorm:"column:user_share;type:decimal(65,18);not null;default:0;comment:userShare"`
	Top       string  `gorm:"column:top;type:varchar(42);not null;index:idx_top_time,priority:1;comment:top"`

	Pool             float64 `gorm:"column:pool;type:decimal(65,18);not null;default:0;comment:pool"`
	UplinePortionBps uint64  `gorm:"column:upline_portion_bps;type:bigint unsigned;not null;default:0;comment:uplinePortionBps"`

	ToL1      float64 `gorm:"column:to_l1;type:decimal(65,18);not null;default:0;comment:toL1"`
	ToL2      float64 `gorm:"column:to_l2;type:decimal(65,18);not null;default:0;comment:toL2"`
	ToTop     float64 `gorm:"column:to_top;type:decimal(65,18);not null;default:0;comment:toTop"`
	ToProject float64 `gorm:"column:to_project;type:decimal(65,18);not null;default:0;comment:toProject"`

	CreatedAt time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;comment:更新时间"`
}

type RewardDetail struct {
	ID         uint64    `gorm:"primarykey;type:bigint unsigned;comment:主键"`
	User       string    `gorm:"column:user;type:varchar(42);not null;index:idx_user_time,priority:1;comment:indexed user"`
	Amount     float64   `gorm:"column:amount;type:decimal(65,18);not null;default:0;comment:amount"`
	Reason     uint64    `gorm:"column:reason;type:bigint unsigned;not null;default:0;comment:reason"`
	NotifiedId uint64    `gorm:"column:notified_id;type:int unsigned;not null;default:0;comment:notified_id"`
	CreatedAt  time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null;comment:更新时间"`
	BlockTime  uint64    `gorm:"column:block_time;type:bigint unsigned;not null;index:idx_block_time;comment:区块时间(秒)"`
}

// data/nft_market_purchase.go
type NftMarketPurchase struct {
	ID uint64 `gorm:"primaryKey;autoIncrement;column:id"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null"`

	Buyer   string `gorm:"column:buyer;type:varchar(42);not null"`
	Seller  string `gorm:"column:seller;type:varchar(42);not null"`
	TokenID uint64 `gorm:"column:token_id;not null"`

	PriceUSDT  float64 `gorm:"column:price_usdt;type:decimal(65,18);not null"`
	FeePaidInB uint8   `gorm:"column:fee_paid_in_b;not null"`
	FeeUSDT    float64 `gorm:"column:fee_usdt;type:decimal(65,18);not null"`
	FeeB       float64 `gorm:"column:fee_b;type:decimal(65,18);not null"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"` // 你表没有也无所谓
	CheckStatus uint64    `gorm:"not null"`
}

type NftMinted struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`

	BlockNumber uint64 `gorm:"column:block_number;not null" json:"block_number"`
	BlockTime   uint64 `gorm:"column:block_time;not null" json:"block_time"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0" json:"log_index"`

	ToAddr  string `gorm:"column:to_addr;type:varchar(42);not null" json:"to_addr"`
	TokenID uint64 `gorm:"column:token_id;not null" json:"token_id"`

	Tier     uint64  `gorm:"column:tier;not null;default:0" json:"tier"`
	UsdtPaid float64 `gorm:"column:usdt_paid;not null;default:0" json:"usdt_paid"`

	Status   uint8  `gorm:"column:status;not null;default:0" json:"status"`
	ListedAt uint64 `gorm:"column:listed_at;not null;default:0" json:"listed_at"`

	OpenStatus uint8  `gorm:"column:open_status;not null;default:0" json:"open_status"`
	OpenedAt   uint64 `gorm:"column:opened_at;not null;default:0" json:"opened_at"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	CheckStatus uint64    `gorm:"not null"`
	CheckTime   uint64    `gorm:"not null"`

	MintAddr string `gorm:"column:mint_addr;type:varchar(42);not null" json:"mint_addr"`
}

type NftMarketListed struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	Seller    string `gorm:"column:seller;type:varchar(42);not null"`
	TokenID   uint64 `gorm:"column:token_id;not null"`
	Timestamp uint64 `gorm:"column:timestamp;not null;default:0"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CheckStatus uint64    `gorm:"not null"`
}

type NftMarketUnlisted struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	Operator string `gorm:"column:operator;type:varchar(42);not null"`
	TokenID  uint64 `gorm:"column:token_id;not null"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CheckStatus uint64    `gorm:"not null"`
}

type NftOpened struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	UserAddr string `gorm:"column:user_addr;type:varchar(42);not null"`
	TokenID  uint64 `gorm:"column:token_id;not null" json:"token_id"`

	OpenedAt uint64  `gorm:"column:opened_at;not null;default:0"`
	Reward   float64 `gorm:"column:reward;type:decimal(65,18);not null;default:0"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CheckStatus uint64    `gorm:"not null"`
}

type NftTransfer struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	FromAddr string `gorm:"column:from_addr;type:varchar(42);not null"`
	ToAddr   string `gorm:"column:to_addr;type:varchar(42);not null"`
	TokenID  uint64 `gorm:"column:token_id;not null"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CheckStatus uint64    `gorm:"not null"`
}

type UserRegistered struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	UserAddr   string `gorm:"column:user_addr;type:varchar(42);not null"`
	ParentAddr string `gorm:"column:parent_addr;type:varchar(42);not null"`
	TopAddr    string `gorm:"column:top_addr;type:varchar(42);not null"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type BindReferral struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`

	BlockNumber uint64 `gorm:"column:block_number;not null" json:"block_number"`
	BlockTime   uint64 `gorm:"column:block_time;not null" json:"block_time"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0" json:"log_index"`

	UserAddr   string `gorm:"column:user_addr;type:varchar(42);not null" json:"user_addr"`
	ParentAddr string `gorm:"column:parent_addr;type:varchar(42);not null" json:"parent_addr"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	CheckStatus uint64 `gorm:"column:check_status;not null;default:0" json:"check_status"`
	CheckTime   uint64 `gorm:"column:check_time;not null;default:0" json:"check_time"`
	Level       int
}

type UserV1Bound struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`

	BlockNumber uint64 `gorm:"column:block_number;not null" json:"block_number"`

	UserAddr      string `gorm:"column:user_addr;type:varchar(42);not null" json:"user_addr"`
	ParentAddr    string `gorm:"column:parent_addr;type:varchar(42);not null" json:"parent_addr"`
	RecommendCode string `gorm:"column:recommend_code;type:varchar(4096);not null;default:''" json:"recommend_code"`

	Amount                            string `gorm:"column:amount;type:decimal(65,18);not null;default:0"`
	AmountHistory                     string `gorm:"column:amount_history;type:decimal(65,18);not null;default:0"`
	ChildrenAmount                    string `gorm:"column:children_amount;type:decimal(65,18);not null;default:0"`
	ChildrenAmountHistory             string `gorm:"column:children_amount_history;type:decimal(65,18);not null;default:0"`
	ChildrenAmountExtra               string `gorm:"column:children_amount_extra;type:decimal(65,18);not null;default:0"`
	RewardRecommendAmount             string `gorm:"column:reward_recommend_amount;type:decimal(65,18);not null;default:0"`
	RewardRecommendPay                string `gorm:"column:reward_recommend_pay;type:decimal(65,18);not null;default:0"`
	RewardRecommendStoreAmount        string `gorm:"column:reward_recommend_store_amount;type:decimal(65,18);not null;default:0"`
	RewardRecommendFee                string `gorm:"column:reward_recommend_fee;type:decimal(65,18);not null;default:0"`
	RewardRecommendTeamUAmount        string `gorm:"column:reward_recommend_team_u_amount;type:decimal(65,18);not null;default:0"`
	RewardRecommendClaimedTeamUNet    string `gorm:"column:reward_recommend_claimed_team_u_net;type:decimal(65,18);not null;default:0"`
	RewardRecommendClaimedTeamUAmount string `gorm:"column:reward_recommend_claimed_team_u_amount;type:decimal(65,18);not null;default:0"`
	RewardRecommendClaimedTeamUFee    string `gorm:"column:reward_recommend_claimed_team_u_fee;type:decimal(65,18);not null;default:0"`
	RewardRecommendExpired            string `gorm:"column:reward_recommend_expired;type:decimal(65,18);not null;default:0"`
	LineU                             string `gorm:"column:line_u;type:decimal(65,18);not null;default:0"`
	LineCoinU                         string `gorm:"column:line_coin_u;type:decimal(65,18);not null;default:0"`
	LineCoin                          string `gorm:"column:line_coin;type:decimal(65,18);not null;default:0"`
	LineFee                           string `gorm:"column:line_fee;type:decimal(65,18);not null;default:0"`
	LevelReward                       string `gorm:"column:level_reward;type:decimal(65,18);not null;default:0"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

type UserV1BoundSyncProgress struct {
	ID                 uint8     `gorm:"column:id;primaryKey"`
	LastProcessedBlock uint64    `gorm:"column:last_processed_block;not null"`
	CreatedAt          time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt          time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type UserV1PerformanceSyncProgress struct {
	StreamName         string    `gorm:"column:stream_name;primaryKey;type:varchar(32)"`
	LastProcessedBlock uint64    `gorm:"column:last_processed_block;not null"`
	CreatedAt          time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt          time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type UserV1StakeChanged struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	BlockNumber uint64    `gorm:"column:block_number;not null"`
	EventKey    string    `gorm:"column:event_key;type:varchar(96);not null;uniqueIndex"`
	TxHash      string    `gorm:"column:tx_hash;type:varchar(66);not null"`
	UserAddr    string    `gorm:"column:user_addr;type:varchar(42);not null"`
	Amount      string    `gorm:"column:amount;type:decimal(65,18);not null;default:0"`
	IsAdd       bool      `gorm:"column:is_add;not null;default:0"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type UserV1ExtraChanged struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	BlockNumber uint64    `gorm:"column:block_number;not null"`
	EventKey    string    `gorm:"column:event_key;type:varchar(96);not null;uniqueIndex"`
	TxHash      string    `gorm:"column:tx_hash;type:varchar(66);not null"`
	UserAddr    string    `gorm:"column:user_addr;type:varchar(42);not null"`
	ExtraAmount string    `gorm:"column:extra_amount;type:decimal(65,18);not null;default:0"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type StakingV1TeamBooked struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	BlockNumber uint64    `gorm:"column:block_number;not null"`
	EventKey    string    `gorm:"column:event_key;type:varchar(96);not null;uniqueIndex"`
	TxHash      string    `gorm:"column:tx_hash;type:varchar(66);not null"`
	FromAddr    string    `gorm:"column:from_addr;type:varchar(42);not null"`
	ToAddr      string    `gorm:"column:to_addr;type:varchar(42);not null"`
	Amount      string    `gorm:"column:amount;type:decimal(65,18);not null;default:0"`
	StoreAmount string    `gorm:"column:store_amount;type:decimal(65,18);not null;default:0"`
	Pay         string    `gorm:"column:pay;type:decimal(65,18);not null;default:0"`
	Fee         string    `gorm:"column:fee;type:decimal(65,18);not null;default:0"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type StakingV1TeamClaimed struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	BlockNumber uint64    `gorm:"column:block_number;not null"`
	EventKey    string    `gorm:"column:event_key;type:varchar(96);not null;uniqueIndex"`
	TxHash      string    `gorm:"column:tx_hash;type:varchar(66);not null"`
	UserAddr    string    `gorm:"column:user_addr;type:varchar(42);not null"`
	Amount      string    `gorm:"column:amount;type:decimal(65,18);not null;default:0"`
	Fee         string    `gorm:"column:fee;type:decimal(65,18);not null;default:0"`
	Net         string    `gorm:"column:net;type:decimal(65,18);not null;default:0"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type StakingV1TeamExpired struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	BlockNumber uint64    `gorm:"column:block_number;not null"`
	EventKey    string    `gorm:"column:event_key;type:varchar(96);not null;uniqueIndex"`
	TxHash      string    `gorm:"column:tx_hash;type:varchar(66);not null"`
	FromAddr    string    `gorm:"column:from_addr;type:varchar(42);not null"`
	ToAddr      string    `gorm:"column:to_addr;type:varchar(42);not null"`
	Amount      string    `gorm:"column:amount;type:decimal(65,18);not null;default:0"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type StakingV1LineClaimed struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	BlockNumber uint64    `gorm:"column:block_number;not null"`
	EventKey    string    `gorm:"column:event_key;type:varchar(96);not null;uniqueIndex"`
	TxHash      string    `gorm:"column:tx_hash;type:varchar(66);not null"`
	UserAddr    string    `gorm:"column:user_addr;type:varchar(42);not null"`
	OrderID     string    `gorm:"column:order_id;type:decimal(65,0);not null;default:0"`
	GrossU      string    `gorm:"column:gross_u;type:decimal(65,18);not null;default:0"`
	FeeU        string    `gorm:"column:fee_u;type:decimal(65,18);not null;default:0"`
	PaidMs      bool      `gorm:"column:paid_ms;not null;default:0"`
	MsAmount    string    `gorm:"column:ms_amount;type:decimal(65,18);not null;default:0"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type StakingStaked struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`

	BlockNumber uint64 `gorm:"column:block_number;not null" json:"block_number"`
	BlockTime   uint64 `gorm:"column:block_time;not null" json:"block_time"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0" json:"log_index"`

	UserAddr   string  `gorm:"column:user_addr;type:varchar(42);not null" json:"user_addr"`
	Amount     float64 `gorm:"column:amount;type:decimal(65,18);not null;default:0" json:"amount"`
	Timestamp  uint64  `gorm:"column:timestamp;not null;default:0" json:"timestamp"`
	StakeIndex uint64  `gorm:"column:stake_index;not null;default:0" json:"stake_index"`
	Duration   uint64  `gorm:"column:duration;not null;default:0" json:"duration"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	CheckStatus uint64 `gorm:"column:check_status;not null;default:0" json:"check_status"`
	CheckTime   uint64 `gorm:"column:check_time;not null;default:0" json:"check_time"`
}

type StakingUnstaked struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`

	BlockNumber uint64 `gorm:"column:block_number;not null" json:"block_number"`
	BlockTime   uint64 `gorm:"column:block_time;not null" json:"block_time"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0" json:"log_index"`

	UserAddr   string  `gorm:"column:user_addr;type:varchar(42);not null" json:"user_addr"`
	Amount     float64 `gorm:"column:amount;type:decimal(65,18);not null;default:0" json:"amount"`
	Timestamp  uint64  `gorm:"column:timestamp;not null;default:0" json:"timestamp"`
	StakeIndex uint64  `gorm:"column:stake_index;not null;default:0" json:"stake_index"`
	Reward     float64 `gorm:"column:reward;type:decimal(65,18);not null;default:0" json:"reward"`
	TTL        uint64  `gorm:"column:ttl;not null;default:0" json:"ttl"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	CheckStatus uint64 `gorm:"column:check_status;not null;default:0" json:"check_status"`
	CheckTime   uint64 `gorm:"column:check_time;not null;default:0" json:"check_time"`
}

type StakingQueueAdded struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`

	BlockNumber uint64 `gorm:"column:block_number;not null" json:"block_number"`
	BlockTime   uint64 `gorm:"column:block_time;not null" json:"block_time"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0" json:"log_index"`

	QueueIndex uint64  `gorm:"column:queue_index;not null;default:0" json:"queue_index"`
	UserAddr   string  `gorm:"column:user_addr;type:varchar(42);not null" json:"user_addr"`
	Amount     float64 `gorm:"column:amount;type:decimal(65,18);not null;default:0" json:"amount"`
	StakeIndex uint8   `gorm:"column:stake_index;not null;default:0" json:"stake_index"`
	QueuedAt   uint64  `gorm:"column:queued_at;not null;default:0" json:"queued_at"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	CheckStatus uint64 `gorm:"column:check_status;not null;default:0" json:"check_status"`
	CheckTime   uint64 `gorm:"column:check_time;not null;default:0" json:"check_time"`
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *UserRepo) GetSwapTradeLast(ctx context.Context) (*biz.SwapTrade, error) {
	var v *SwapTrade

	if err := u.data.DB(ctx).Table("swap_trade").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "SWAP_TRADE_ERROR", err.Error())
	}

	return &biz.SwapTrade{
		ID:              v.ID,
		BlockNumber:     v.BlockNumber,
		LogIndex:        v.LogIndex,
		BlockTime:       v.BlockTime,
		Sender:          v.Sender,
		ToAddr:          v.ToAddr,
		Side:            v.Side,
		Amount0In:       v.Amount0In,
		Amount1In:       v.Amount1In,
		Amount0OutNet:   v.Amount0OutNet,
		Amount1OutGross: v.Amount1OutGross,
		Amount0OutGross: v.Amount0OutGross,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetSwapTrade(ctx context.Context, start, end uint64) ([]*biz.SwapTrade, error) {
	var s []*SwapTrade

	res := make([]*biz.SwapTrade, 0)
	if err := u.data.DB(ctx).Table("swap_trade").Where("block_time >=?", start).Where("block_time <=?", end).Find(&s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "SWAP_TRADE_ERROR", err.Error())
	}

	for _, v := range s {
		res = append(res, &biz.SwapTrade{
			ID:              v.ID,
			BlockNumber:     v.BlockNumber,
			LogIndex:        v.LogIndex,
			BlockTime:       v.BlockTime,
			Sender:          v.Sender,
			ToAddr:          v.ToAddr,
			Side:            v.Side,
			Amount0In:       v.Amount0In,
			Amount1In:       v.Amount1In,
			Amount0OutNet:   v.Amount0OutNet,
			Amount1OutGross: v.Amount1OutGross,
			Amount0OutGross: v.Amount0OutGross,
			CreatedAt:       v.CreatedAt,
			UpdatedAt:       v.UpdatedAt,
		})
	}
	return res, nil
}

// InsertSwapTrade .
func (u *UserRepo) InsertSwapTrade(ctx context.Context, iData *biz.SwapTrade) error {
	var (
		err error
		s   SwapTrade
	)

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.Side = iData.Side
	s.Sender = iData.Sender
	s.Amount1In = iData.Amount1In
	s.Amount0OutGross = iData.Amount0OutGross
	s.Amount0OutNet = iData.Amount0OutNet
	s.Amount1OutGross = iData.Amount1OutGross
	s.Amount0In = iData.Amount0In
	s.ToAddr = iData.ToAddr
	s.LogIndex = iData.LogIndex

	err = u.data.DB(ctx).Table("swap_trade").Create(&s).Error
	if err != nil {
		return errors.New(500, "CREATE_SWAP_TRADE_ERROR", "信息创建失败")
	}

	return nil
}

func (u *UserRepo) GetPrimaryBuyLast(ctx context.Context) (*biz.PrimaryBuy, error) {
	var v PrimaryBuy

	if err := u.data.DB(ctx).Table("primary_buy").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "PRIMARY_BUY_ERROR", err.Error())
	}

	return &biz.PrimaryBuy{
		ID:           v.ID,
		BlockNumber:  v.BlockNumber,
		BlockTime:    v.BlockTime,
		LogIndex:     v.LogIndex,
		Buyer:        v.Buyer,
		ToAddr:       v.ToAddr,
		UsdtUsed:     v.UsdtUsed,
		AusdGrossOut: v.AusdGrossOut,
		AusdFee:      v.AusdFee,
		AusdNetOut:   v.AusdNetOut,
		PriceBefore:  v.PriceBefore,
		PriceAfter:   v.PriceAfter,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetPrimaryBuy(ctx context.Context, start, end uint64) ([]*biz.PrimaryBuy, error) {
	var s []PrimaryBuy
	res := make([]*biz.PrimaryBuy, 0)

	if err := u.data.DB(ctx).
		Table("primary_buy").
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Order("block_time asc, id asc").
		Find(&s).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "PRIMARY_BUY_ERROR", err.Error())
	}

	for _, v := range s {
		res = append(res, &biz.PrimaryBuy{
			ID:           v.ID,
			BlockNumber:  v.BlockNumber,
			BlockTime:    v.BlockTime,
			LogIndex:     v.LogIndex,
			Buyer:        v.Buyer,
			ToAddr:       v.ToAddr,
			UsdtUsed:     v.UsdtUsed,
			AusdGrossOut: v.AusdGrossOut,
			AusdFee:      v.AusdFee,
			AusdNetOut:   v.AusdNetOut,
			PriceBefore:  v.PriceBefore,
			PriceAfter:   v.PriceAfter,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}
	return res, nil
}

// InsertPrimaryBuy .
func (u *UserRepo) InsertPrimaryBuy(ctx context.Context, iData *biz.PrimaryBuy) error {
	var s PrimaryBuy

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex
	s.Buyer = iData.Buyer
	s.ToAddr = iData.ToAddr

	s.UsdtUsed = iData.UsdtUsed
	s.AusdGrossOut = iData.AusdGrossOut
	s.AusdFee = iData.AusdFee
	s.AusdNetOut = iData.AusdNetOut
	s.PriceBefore = iData.PriceBefore
	s.PriceAfter = iData.PriceAfter

	if err := u.data.DB(ctx).Table("primary_buy").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_PRIMARY_BUY_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetPrimarySellLast(ctx context.Context) (*biz.PrimarySell, error) {
	var v PrimarySell

	if err := u.data.DB(ctx).Table("primary_sell").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "PRIMARY_SELL_ERROR", err.Error())
	}

	return &biz.PrimarySell{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,
		Seller:      v.Seller,
		ToAddr:      v.ToAddr,
		AusdGrossIn: v.AusdGrossIn,
		AusdFee:     v.AusdFee,
		AusdBurn:    v.AusdBurn,
		UsdtOut:     v.UsdtOut,
		PriceBefore: v.PriceBefore,
		PriceAfter:  v.PriceAfter,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetPrimarySell(ctx context.Context, start, end uint64) ([]*biz.PrimarySell, error) {
	var s []PrimarySell
	res := make([]*biz.PrimarySell, 0)

	if err := u.data.DB(ctx).
		Table("primary_sell").
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Order("block_time asc, id asc").
		Find(&s).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "PRIMARY_SELL_ERROR", err.Error())
	}

	for _, v := range s {
		res = append(res, &biz.PrimarySell{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,
			Seller:      v.Seller,
			ToAddr:      v.ToAddr,
			AusdGrossIn: v.AusdGrossIn,
			AusdFee:     v.AusdFee,
			AusdBurn:    v.AusdBurn,
			UsdtOut:     v.UsdtOut,
			PriceBefore: v.PriceBefore,
			PriceAfter:  v.PriceAfter,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	return res, nil
}

// InsertPrimarySell .
func (u *UserRepo) InsertPrimarySell(ctx context.Context, iData *biz.PrimarySell) error {
	var s PrimarySell

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex
	s.Seller = iData.Seller
	s.ToAddr = iData.ToAddr

	s.AusdGrossIn = iData.AusdGrossIn
	s.AusdFee = iData.AusdFee
	s.AusdBurn = iData.AusdBurn
	s.UsdtOut = iData.UsdtOut
	s.PriceBefore = iData.PriceBefore
	s.PriceAfter = iData.PriceAfter

	if err := u.data.DB(ctx).Table("primary_sell").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_PRIMARY_SELL_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetRewardNotifiedLast(ctx context.Context) (*biz.RewardNotified, error) {
	var v RewardNotified

	if err := u.data.DB(ctx).Table("reward_notified").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "REWARD_NOTIFIED_ERROR", err.Error())
	}

	return &biz.RewardNotified{
		ID:               v.ID,
		BlockNumber:      v.BlockNumber,
		BlockTime:        v.BlockTime,
		LogIndex:         v.LogIndex,
		User:             v.User,
		L1:               v.L1,
		L2:               v.L2,
		Profit:           v.Profit,
		UserShare:        v.UserShare,
		Top:              v.Top,
		Pool:             v.Pool,
		UplinePortionBps: v.UplinePortionBps,
		ToL1:             v.ToL1,
		ToL2:             v.ToL2,
		ToTop:            v.ToTop,
		ToProject:        v.ToProject,
		CreatedAt:        v.CreatedAt,
		UpdatedAt:        v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetRewardNotified(ctx context.Context, start, end uint64) ([]*biz.RewardNotified, error) {
	var s []RewardNotified
	res := make([]*biz.RewardNotified, 0)

	if err := u.data.DB(ctx).
		Table("reward_notified").
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Order("block_time asc, id asc").
		Find(&s).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "REWARD_NOTIFIED_ERROR", err.Error())
	}

	for _, v := range s {
		res = append(res, &biz.RewardNotified{
			ID:               v.ID,
			BlockNumber:      v.BlockNumber,
			BlockTime:        v.BlockTime,
			LogIndex:         v.LogIndex,
			User:             v.User,
			L1:               v.L1,
			L2:               v.L2,
			Profit:           v.Profit,
			UserShare:        v.UserShare,
			Top:              v.Top,
			Pool:             v.Pool,
			UplinePortionBps: v.UplinePortionBps,
			ToL1:             v.ToL1,
			ToL2:             v.ToL2,
			ToTop:            v.ToTop,
			ToProject:        v.ToProject,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		})
	}
	return res, nil
}

func (u *UserRepo) GetRewardNotifiedByIds(ctx context.Context, ids []uint64) (map[uint64]*biz.RewardNotified, error) {
	var s []RewardNotified
	res := make(map[uint64]*biz.RewardNotified, 0)

	if err := u.data.DB(ctx).
		Table("reward_notified").
		Where("id in(?)", ids).
		Find(&s).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "REWARD_NOTIFIED_ERROR", err.Error())
	}

	for _, v := range s {
		res[v.ID] = &biz.RewardNotified{
			ID:               v.ID,
			BlockNumber:      v.BlockNumber,
			BlockTime:        v.BlockTime,
			LogIndex:         v.LogIndex,
			User:             v.User,
			L1:               v.L1,
			L2:               v.L2,
			Profit:           v.Profit,
			UserShare:        v.UserShare,
			Top:              v.Top,
			Pool:             v.Pool,
			UplinePortionBps: v.UplinePortionBps,
			ToL1:             v.ToL1,
			ToL2:             v.ToL2,
			ToTop:            v.ToTop,
			ToProject:        v.ToProject,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		}
	}
	return res, nil
}

func (u *UserRepo) InsertRewardNotified(ctx context.Context, iData *biz.RewardNotified) error {
	var s RewardNotified

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.User = iData.User
	s.L1 = iData.L1
	s.L2 = iData.L2

	s.Profit = iData.Profit
	s.UserShare = iData.UserShare
	s.Top = iData.Top

	s.Pool = iData.Pool
	s.UplinePortionBps = iData.UplinePortionBps

	s.ToL1 = iData.ToL1
	s.ToL2 = iData.ToL2
	s.ToTop = iData.ToTop
	s.ToProject = iData.ToProject

	if err := u.data.DB(ctx).Table("reward_notified").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_NOTIFIED_ERROR", "信息创建失败")
	}

	var one RewardDetail
	one.User = s.User
	one.Reason = 1
	one.NotifiedId = s.ID
	one.Amount = s.UserShare
	one.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&one).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	var two RewardDetail
	two.User = s.L1
	two.Reason = 2
	two.NotifiedId = s.ID
	two.Amount = s.ToL1
	two.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&two).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	var three RewardDetail
	three.User = s.L2
	three.Reason = 3
	three.NotifiedId = s.ID
	three.Amount = s.ToL2
	three.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&three).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	var four RewardDetail
	four.User = s.Top
	four.Reason = 4
	four.NotifiedId = s.ID
	four.Amount = s.ToTop
	four.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&four).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	var five RewardDetail
	five.User = "0x144351b1a5Af4538f53556633D8f10495aA564A8"
	five.Reason = 5
	five.NotifiedId = s.ID
	five.Amount = s.ToProject
	five.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&five).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	return nil
}

// GetUserRewardByUserIdPage .
func (u *UserRepo) GetUserRewardByUserIdPage(ctx context.Context, b *biz.Pagination, address string, reason uint64) ([]*biz.RewardDetail, error, int64) {
	var (
		count   int64
		rewards []*RewardDetail
	)

	res := make([]*biz.RewardDetail, 0)

	instance := u.data.db.Where("user", address).Table("reward_detail").Order("id desc")
	if 0 < reason {
		instance = instance.Where("reason=?", reason)
	}

	instance = instance.Count(&count)

	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found"), 0
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error()), 0
	}

	for _, reward := range rewards {
		res = append(res, &biz.RewardDetail{
			ID:         reward.ID,
			User:       reward.User,
			Amount:     reward.Amount,
			Reason:     reward.Reason,
			BlockTime:  reward.BlockTime,
			NotifiedId: reward.NotifiedId,
			CreatedAt:  reward.CreatedAt,
			UpdatedAt:  reward.UpdatedAt,
		})
	}

	return res, nil, count
}

// GetNftMarketPurchaseByAddressPage
// side: 0=buyer或seller  1=buyer  2=seller
func (u *UserRepo) GetNftMarketPurchaseByAddressPage(
	ctx context.Context,
	b *biz.Pagination,
	address string,
	addressTwo []string,
	side uint64,
) ([]*biz.NftMarketPurchase, error, int64) {

	var (
		count int64
		rows  []*NftMarketPurchase
	)

	res := make([]*biz.NftMarketPurchase, 0)

	instance := u.data.DB(ctx).Table("nft_market_purchase").Order("id desc")

	if address != "" {
		if side == 1 {
			instance = instance.Where("buyer = ?", address)
		} else if side == 2 {
			instance = instance.Where("seller = ?", address)
		} else if side == 3 {
			instance = instance.Where("seller in (?)", addressTwo)
		} else {
			instance = instance.Where("(buyer = ? OR seller = ?)", address, address)
		}
	}

	instance = instance.Count(&count)

	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&rows).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("NFT_PURCHASE_NOT_FOUND", "purchase not found"), 0
		}
		return nil, errors.New(500, "NFT_PURCHASE_PAGE_ERROR", err.Error()), 0
	}

	for _, r := range rows {
		res = append(res, &biz.NftMarketPurchase{
			ID:          r.ID,
			BlockNumber: r.BlockNumber,
			BlockTime:   r.BlockTime,
			LogIndex:    r.LogIndex,

			Buyer:   r.Buyer,
			Seller:  r.Seller,
			TokenID: r.TokenID,

			PriceUSDT:  r.PriceUSDT,
			FeePaidInB: r.FeePaidInB,
			FeeUSDT:    r.FeeUSDT,
			FeeB:       r.FeeB,

			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}

	return res, nil, count
}

func (u *UserRepo) InsertNftMarketPurchase(ctx context.Context, iData *biz.NftMarketPurchase) error {
	var s NftMarketPurchase

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.Buyer = iData.Buyer
	s.Seller = iData.Seller
	s.TokenID = iData.TokenID

	s.PriceUSDT = iData.PriceUSDT
	s.FeePaidInB = iData.FeePaidInB
	s.FeeUSDT = iData.FeeUSDT
	s.FeeB = iData.FeeB

	if err := u.data.DB(ctx).Table("nft_market_purchase").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_MARKET_PURCHASE_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftMarketPurchaseLast(ctx context.Context) (*biz.NftMarketPurchase, error) {
	var v NftMarketPurchase

	if err := u.data.DB(ctx).Table("nft_market_purchase").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_MARKET_PURCHASE_ERROR", err.Error())
	}

	return &biz.NftMarketPurchase{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		Buyer:   v.Buyer,
		Seller:  v.Seller,
		TokenID: v.TokenID,

		PriceUSDT:  v.PriceUSDT,
		FeePaidInB: v.FeePaidInB,
		FeeUSDT:    v.FeeUSDT,
		FeeB:       v.FeeB,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) UpdateNftMinted(ctx context.Context, tokenId uint64, mintAddr string) error {
	res := u.data.DB(ctx).Table("nft_minted").Where("token_id=?", tokenId).
		Updates(map[string]interface{}{
			"mint_addr":  mintAddr,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_MT_ERROR", "修改失败")
	}
	return nil
}

func (u *UserRepo) InsertNftMinted(ctx context.Context, iData *biz.NftMinted) error {
	var s NftMinted

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.ToAddr = iData.ToAddr
	s.TokenID = iData.TokenID

	s.Tier = iData.Tier
	s.UsdtPaid = iData.UsdtPaid
	s.MintAddr = iData.MintAddr

	// ✅ 默认值交给 DB：status/open_status/listed_at/opened_at
	// 如果你就是要插入时强行写，也可以从 iData 赋值

	if err := u.data.DB(ctx).Table("nft_minted").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_MINTED_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftMintedLast(ctx context.Context) (*biz.NftMinted, error) {
	var v NftMinted

	if err := u.data.DB(ctx).Table("nft_minted").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_MINTED_ERROR", err.Error())
	}

	return &biz.NftMinted{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		ToAddr:  v.ToAddr,
		TokenID: v.TokenID,

		Tier:     v.Tier,
		UsdtPaid: v.UsdtPaid,

		Status:   v.Status,
		ListedAt: v.ListedAt,

		OpenStatus: v.OpenStatus,
		OpenedAt:   v.OpenedAt,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) InsertNftMarketListed(ctx context.Context, iData *biz.NftMarketListed) error {
	var s NftMarketListed

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.Seller = iData.Seller
	s.TokenID = iData.TokenID
	s.Timestamp = iData.Timestamp

	if err := u.data.DB(ctx).Table("nft_market_listed").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_MARKET_LISTED_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftMarketListedLast(ctx context.Context) (*biz.NftMarketListed, error) {
	var v NftMarketListed

	if err := u.data.DB(ctx).Table("nft_market_listed").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_MARKET_LISTED_ERROR", err.Error())
	}

	return &biz.NftMarketListed{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		Seller:    v.Seller,
		TokenID:   v.TokenID,
		Timestamp: v.Timestamp,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) InsertNftMarketUnlisted(ctx context.Context, iData *biz.NftMarketUnlisted) error {
	var s NftMarketUnlisted

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.Operator = iData.Operator
	s.TokenID = iData.TokenID

	if err := u.data.DB(ctx).Table("nft_market_unlisted").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_MARKET_UNLISTED_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftMarketUnlistedLast(ctx context.Context) (*biz.NftMarketUnlisted, error) {
	var v NftMarketUnlisted

	if err := u.data.DB(ctx).Table("nft_market_unlisted").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_MARKET_UNLISTED_ERROR", err.Error())
	}

	return &biz.NftMarketUnlisted{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		Operator: v.Operator,
		TokenID:  v.TokenID,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) InsertNftOpened(ctx context.Context, iData *biz.NftOpened) error {
	var s NftOpened

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.UserAddr = iData.UserAddr
	s.TokenID = iData.TokenID

	s.OpenedAt = iData.OpenedAt
	s.Reward = iData.Reward

	if err := u.data.DB(ctx).Table("nft_opened").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_OPENED_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftOpenedLast(ctx context.Context) (*biz.NftOpened, error) {
	var v NftOpened

	if err := u.data.DB(ctx).Table("nft_opened").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_OPENED_ERROR", err.Error())
	}

	return &biz.NftOpened{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		UserAddr: v.UserAddr,
		TokenID:  v.TokenID,

		OpenedAt: v.OpenedAt,
		Reward:   v.Reward,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) InsertNftTransfer(ctx context.Context, iData *biz.NftTransfer) error {
	var s NftTransfer

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.FromAddr = iData.FromAddr
	s.ToAddr = iData.ToAddr
	s.TokenID = iData.TokenID

	if err := u.data.DB(ctx).Table("nft_transfer").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_TRANSFER_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftTransferLast(ctx context.Context) (*biz.NftTransfer, error) {
	var v NftTransfer

	if err := u.data.DB(ctx).Table("nft_transfer").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	return &biz.NftTransfer{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		FromAddr: v.FromAddr,
		ToAddr:   v.ToAddr,
		TokenID:  v.TokenID,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetNftTransferLastNoCheck(ctx context.Context) ([]*biz.NftTransfer, error) {
	var vL []*NftTransfer

	res := make([]*biz.NftTransfer, 0)
	if err := u.data.DB(ctx).Table("nft_transfer").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftTransfer{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			FromAddr: v.FromAddr,
			ToAddr:   v.ToAddr,
			TokenID:  v.TokenID,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftListLastNoCheck(ctx context.Context) ([]*biz.NftMarketListed, error) {
	var vL []*NftMarketListed

	res := make([]*biz.NftMarketListed, 0)
	if err := u.data.DB(ctx).Table("nft_market_listed").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftMarketListed{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			Seller:    v.Seller,
			TokenID:   v.TokenID,
			Timestamp: v.Timestamp,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftUnListLastNoCheck(ctx context.Context) ([]*biz.NftMarketUnlisted, error) {
	var vL []*NftMarketUnlisted

	res := make([]*biz.NftMarketUnlisted, 0)
	if err := u.data.DB(ctx).Table("nft_market_unlisted").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftMarketUnlisted{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			Operator: v.Operator,
			TokenID:  v.TokenID,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftBuyLastNoCheck(ctx context.Context) ([]*biz.NftMarketPurchase, error) {
	var vL []*NftMarketPurchase

	res := make([]*biz.NftMarketPurchase, 0)
	if err := u.data.DB(ctx).Table("nft_market_purchase").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftMarketPurchase{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			Buyer:   v.Buyer,
			Seller:  v.Seller,
			TokenID: v.TokenID,

			PriceUSDT:  v.PriceUSDT,
			FeePaidInB: v.FeePaidInB,
			FeeUSDT:    v.FeeUSDT,
			FeeB:       v.FeeB,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftOpenLastNoCheck(ctx context.Context) ([]*biz.NftOpened, error) {
	var vL []*NftOpened

	res := make([]*biz.NftOpened, 0)
	if err := u.data.DB(ctx).Table("nft_opened").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftOpened{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			UserAddr: v.UserAddr,
			TokenID:  v.TokenID,

			OpenedAt: v.OpenedAt,
			Reward:   v.Reward,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftMintedByTokenIds(ctx context.Context, tokenIds []uint64) (map[uint64]*biz.NftMinted, error) {
	var vL []*NftMinted

	res := make(map[uint64]*biz.NftMinted, 0)
	if err := u.data.DB(ctx).Table("nft_minted").Where("token_id in(?)", tokenIds).Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res[v.TokenID] = &biz.NftMinted{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			ToAddr:  v.ToAddr,
			TokenID: v.TokenID,

			Tier:     v.Tier,
			UsdtPaid: v.UsdtPaid,

			Status:   v.Status,
			ListedAt: v.ListedAt,

			OpenStatus: v.OpenStatus,
			OpenedAt:   v.OpenedAt,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			CheckTime: v.CheckTime,
		}
	}

	return res, nil
}

// UpdateNftMintedToAddress .
func (u *UserRepo) UpdateNftMintedToAddress(ctx context.Context, id, idT, check uint64, toAddr string) error {
	res := u.data.DB(ctx).Table("nft_transfer").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_T_ERROR", "修改失败")
	}

	if 1 == check {
		//.Where("check_time<=?", checkTime).
		resTwo := u.data.DB(ctx).Table("nft_minted").Where("id=?", id).
			Updates(map[string]interface{}{
				"to_addr": toAddr,
				//"check_time": checkTime,
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
			return errors.New(500, "UPDATE_MINTED_ERROR", "修改失败")
		}
	}

	return nil
}

// UpdateNftMintedListStatus .
func (u *UserRepo) UpdateNftMintedListStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_listed").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_L_ERROR", "修改失败")
	}

	//.Where("check_time<=?", checkTime).
	resTwo := u.data.DB(ctx).Table("nft_minted").Where("id=?", id).Where("check_time<?", checkTime).
		Updates(map[string]interface{}{
			"check_time": checkTime,
			"status":     1,
			"listed_at":  checkTime,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "UPDATE_MINTED_L_ERROR", "修改失败")
	}

	return nil
}

// UpdateNftMintedUnListStatus .
func (u *UserRepo) UpdateNftMintedUnListStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_unlisted").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_L_ERROR", "修改失败")
	}

	//.Where("check_time<=?", checkTime).
	resTwo := u.data.DB(ctx).Table("nft_minted").Where("id=?", id).Where("check_time<?", checkTime).
		Updates(map[string]interface{}{
			"check_time": checkTime,
			"status":     0,
			"listed_at":  0,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "UPDATE_MINTED_U_L_ERROR", "修改失败")
	}

	return nil
}

func (u *UserRepo) UpdateUnlistedCheckStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_unlisted").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 2,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_CUL_ERROR", "修改失败")
	}

	return nil
}

func (u *UserRepo) UpdateListedCheckStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_listed").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 2,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_CL_ERROR", "修改失败")
	}

	return nil
}

func (u *UserRepo) UpdateBuyCheckStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_purchase").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 2,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_CB_ERROR", "修改失败")
	}

	return nil
}

// UpdateNftMintedBuyStatus .
func (u *UserRepo) UpdateNftMintedBuyStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_purchase").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_L_ERROR", "修改失败")
	}

	//.Where("check_time<=?", checkTime).
	resTwo := u.data.DB(ctx).Table("nft_minted").Where("id=?", id).Where("check_time<?", checkTime).
		Updates(map[string]interface{}{
			"check_time": checkTime,
			"status":     0,
			"listed_at":  0,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "UPDATE_MINTED_B_ERROR", "修改失败")
	}

	return nil
}

// UpdateNftMintedOpenStatus .
func (u *UserRepo) UpdateNftMintedOpenStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_opened").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_L_ERROR", "修改失败")
	}

	//.Where("check_time<=?", checkTime).
	resTwo := u.data.DB(ctx).Table("nft_minted").Where("id=?", id).
		Updates(map[string]interface{}{
			"open_status": 1,
			"opened_at":   checkTime,
			"updated_at":  time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "UPDATE_MINTED_O_ERROR", "修改失败")
	}

	return nil
}

// GetNftMintedPage .
func (u *UserRepo) GetNftMintedPage(
	ctx context.Context,
	b *biz.Pagination,
	order uint64,
	orderTwo uint64,
	tier uint64,
) ([]*biz.NftMinted, error, int64) {

	var (
		count int64
		rows  []*NftMinted
	)

	res := make([]*biz.NftMinted, 0)

	instance := u.data.DB(ctx).Table("nft_minted").Where("status=?", 1).Where("open_status=?", 0)

	if 0 < tier {
		instance = instance.Where("tier=?", tier)
	}

	instance = instance.Count(&count)

	if 0 < order {
		instance = instance.Order("id asc")
	} else {
		instance = instance.Order("id desc")
	}

	if 0 < orderTwo {
		instance = instance.Order("tier asc")
	} else {
		instance = instance.Order("tier desc")
	}

	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&rows).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("NFT_MINTED_NOT_FOUND", "minted not found"), 0
		}
		return nil, errors.New(500, "NFT_MINTED_ERROR", err.Error()), 0
	}

	for _, v := range rows {
		res = append(res, &biz.NftMinted{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			ToAddr:  v.ToAddr,
			TokenID: v.TokenID,

			Tier:     v.Tier,
			UsdtPaid: v.UsdtPaid,

			Status:   v.Status,
			ListedAt: v.ListedAt,

			OpenStatus: v.OpenStatus,
			OpenedAt:   v.OpenedAt,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			MintAddr:  v.MintAddr,
		})
	}

	return res, nil, count
}

// GetNftMintedByAddressPage .
func (u *UserRepo) GetNftMintedByAddressPage(
	ctx context.Context,
	b *biz.Pagination,
	start uint64,
	end uint64,
	address []string,
	status uint64,
	order uint64,
	tier uint64,
	openStatus uint64,
	openAtOrder uint64,
) ([]*biz.NftMinted, error, int64) {

	var (
		count int64
		rows  []*NftMinted
	)

	res := make([]*biz.NftMinted, 0)

	instance := u.data.DB(ctx).Table("nft_minted").Where("to_addr in(?)", address)

	if 0 < start && 0 < end {
		instance = instance.Where("block_time >= ?", start).Where("block_time <= ?", end)
	}

	if 2 > openStatus {
		if 1 == openStatus {
			instance = instance.Where("open_status=?", 1)
		} else {
			instance = instance.Where("open_status=?", 0)
		}
	}

	if 2 > status {
		if 1 == status {
			instance = instance.Where("status=?", 1)
		} else {
			instance = instance.Where("status=?", 0)
		}
	}

	if 0 < tier {
		instance = instance.Where("tier=?", tier)
	}

	instance = instance.Count(&count)

	if 0 < order {
		instance = instance.Order("id asc")
	} else {
		instance = instance.Order("id desc")
	}

	if 0 < openAtOrder {
		if 1 < openAtOrder {
			instance = instance.Order("opened_at asc")
		} else {
			instance = instance.Order("opened_at desc")
		}
	}

	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&rows).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("NFT_MINTED_NOT_FOUND", "minted not found"), 0
		}
		return nil, errors.New(500, "NFT_MINTED_ERROR", err.Error()), 0
	}

	for _, v := range rows {
		res = append(res, &biz.NftMinted{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			ToAddr:  v.ToAddr,
			TokenID: v.TokenID,

			Tier:     v.Tier,
			UsdtPaid: v.UsdtPaid,

			Status:   v.Status,
			ListedAt: v.ListedAt,

			OpenStatus: v.OpenStatus,
			OpenedAt:   v.OpenedAt,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			MintAddr:  v.MintAddr,
		})
	}

	return res, nil, count
}

func (u *UserRepo) InsertUserRegistered(ctx context.Context, iData *biz.UserRegistered) error {
	var s UserRegistered

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.UserAddr = iData.UserAddr
	s.ParentAddr = iData.ParentAddr
	s.TopAddr = iData.TopAddr

	if err := u.data.DB(ctx).Table("user_registered").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_USER_REGISTERED_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetUserRegisteredLast(ctx context.Context) (*biz.UserRegistered, error) {
	var v UserRegistered

	if err := u.data.DB(ctx).Table("user_registered").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "USER_REGISTERED_ERROR", err.Error())
	}

	return &biz.UserRegistered{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		UserAddr:   v.UserAddr,
		ParentAddr: v.ParentAddr,
		TopAddr:    v.TopAddr,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

// GetUserRCount .
func (u *UserRepo) GetUserRCount(ctx context.Context) int64 {
	var (
		count int64
	)

	u.data.db.Table("user_registered").Count(&count)
	return count
}

// GetUserRCountBySe .
func (u *UserRepo) GetUserRCountBySe(ctx context.Context, start, end uint64) int64 {
	var (
		count int64
	)

	u.data.db.Table("user_registered").Where("block_time >= ?", start).Where("block_time <= ?", end).Count(&count)
	return count
}

// GetMintNftCountBySe .
func (u *UserRepo) GetMintNftCountBySe(ctx context.Context, start, end uint64) int64 {
	var (
		count int64
	)

	u.data.db.Table("nft_minted").Where("block_time >= ?", start).Where("block_time <= ?", end).Count(&count)
	return count
}

// GetMintNftUsdtPaidSumBySe 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetMintNftUsdtPaidSumBySe(ctx context.Context, start, end uint64) string {
	var sum string
	u.data.db.Table("nft_minted").
		Select("COALESCE(SUM(usdt_paid), 0)").
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Scan(&sum)
	return sum
}

// GetMintNftCount .
func (u *UserRepo) GetMintNftCount(paidType uint64) int64 {
	var (
		count int64
	)

	instance := u.data.db.Table("nft_minted")
	if 0 < paidType {
		instance = instance.Where("tier", paidType)
	}
	instance.Count(&count)
	return count
}

// GetMintNftNotOpenCount .
func (u *UserRepo) GetMintNftNotOpenCount(paidType uint64) int64 {
	var (
		count int64
	)

	instance := u.data.db.Table("nft_minted").Where("open_status=?", 0)
	if 0 < paidType {
		instance = instance.Where("tier", paidType)
	}
	instance.Count(&count)
	return count
}

// GetMintNftUsdtPaidSum 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetMintNftUsdtPaidSum(paidType uint64) string {
	var sum string
	instance := u.data.db.Table("nft_minted")
	if 0 < paidType {
		instance = instance.Where("tier", paidType)
	}

	instance.Select("COALESCE(SUM(usdt_paid), 0)").Scan(&sum)
	return sum
}

// GetMintNftNotOpenUsdtPaidSum 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetMintNftNotOpenUsdtPaidSum(paidType uint64) string {
	var sum string
	instance := u.data.db.Table("nft_minted").Where("open_status=?", 0)
	if 0 < paidType {
		instance = instance.Where("tier", paidType)
	}

	instance.Select("COALESCE(SUM(usdt_paid), 0)").Scan(&sum)
	return sum
}

// GetNftBuyCountBySe .
func (u *UserRepo) GetNftBuyCountBySe(ctx context.Context, start, end uint64) int64 {
	var (
		count int64
	)

	u.data.db.Table("nft_market_purchase").Where("block_time >= ?", start).Where("block_time <= ?", end).Count(&count)
	return count
}

// GetNftBuySumBySe 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetNftBuySumBySe(ctx context.Context, start, end uint64) string {
	var sum string
	u.data.db.Table("nft_market_purchase").
		Select("COALESCE(SUM(price_usdt), 0)").
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Scan(&sum)
	return sum
}

// GetNftBuyCount .
func (u *UserRepo) GetNftBuyCount() int64 {
	var (
		count int64
	)

	u.data.db.Table("nft_market_purchase").Count(&count)
	return count
}

// GetNftBuySum 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetNftBuySum() string {
	var sum string
	u.data.db.Table("nft_market_purchase").
		Select("COALESCE(SUM(price_usdt), 0)").
		Scan(&sum)
	return sum
}

// GetNftOpenCountBySe .
func (u *UserRepo) GetNftOpenCountBySe(ctx context.Context, start, end uint64) int64 {
	var (
		count int64
	)

	u.data.db.Table("nft_opened").Where("block_time >= ?", start).Where("block_time <= ?", end).Count(&count)
	return count
}

func (u *UserRepo) GetNftOpenSum() string {
	var sum string
	u.data.db.Table("nft_opened").
		Select("COALESCE(SUM(reward), 0)").
		Scan(&sum)
	return sum
}

func (u *UserRepo) GetBindReferralLast(ctx context.Context) (*biz.BindReferral, error) {
	var v BindReferral

	if err := u.data.DB(ctx).Table("user_bind_referral").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "BIND_REFERRAL_ERROR", err.Error())
	}

	return &biz.BindReferral{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		UserAddr:   v.UserAddr,
		ParentAddr: v.ParentAddr,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,

		CheckStatus: v.CheckStatus,
		CheckTime:   v.CheckTime,
	}, nil
}

func (u *UserRepo) GetBindReferrals(ctx context.Context) ([]*biz.BindReferral, error) {
	var user []*BindReferral

	res := make([]*biz.BindReferral, 0)
	if err := u.data.DB(ctx).Table("user_bind_referral").Order("id asc").Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "BIND_REFERRAL_ERROR", err.Error())
	}

	for _, v := range user {
		res = append(res, &biz.BindReferral{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			UserAddr:   v.UserAddr,
			ParentAddr: v.ParentAddr,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,

			CheckStatus: v.CheckStatus,
			CheckTime:   v.CheckTime,
			Level:       int8(v.Level),
		})
	}
	return res, nil
}

func (u *UserRepo) InsertBindReferral(ctx context.Context, iData *biz.BindReferral) error {
	var s BindReferral

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.UserAddr = iData.UserAddr
	s.ParentAddr = iData.ParentAddr
	s.Level = int(iData.Level)

	// ✅ check_status/check_time 默认值交给 DB
	if err := u.data.DB(ctx).Table("user_bind_referral").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_BIND_REFERRAL_ERROR", "信息创建失败")
	}
	return nil
}

func toBizUserV1Bound(v *UserV1Bound) *biz.UserV1Bound {
	if nil == v {
		return nil
	}

	return &biz.UserV1Bound{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,

		UserAddr:      v.UserAddr,
		ParentAddr:    v.ParentAddr,
		RecommendCode: v.RecommendCode,

		Amount:                            v.Amount,
		AmountHistory:                     v.AmountHistory,
		ChildrenAmount:                    v.ChildrenAmount,
		ChildrenAmountHistory:             v.ChildrenAmountHistory,
		ChildrenAmountExtra:               v.ChildrenAmountExtra,
		RewardRecommendAmount:             v.RewardRecommendAmount,
		RewardRecommendPay:                v.RewardRecommendPay,
		RewardRecommendStoreAmount:        v.RewardRecommendStoreAmount,
		RewardRecommendFee:                v.RewardRecommendFee,
		RewardRecommendTeamUAmount:        v.RewardRecommendTeamUAmount,
		RewardRecommendClaimedTeamUNet:    v.RewardRecommendClaimedTeamUNet,
		RewardRecommendClaimedTeamUAmount: v.RewardRecommendClaimedTeamUAmount,
		RewardRecommendClaimedTeamUFee:    v.RewardRecommendClaimedTeamUFee,
		RewardRecommendExpired:            v.RewardRecommendExpired,
		LineU:                             v.LineU,
		LineCoinU:                         v.LineCoinU,
		LineCoin:                          v.LineCoin,
		LineFee:                           v.LineFee,
		LevelReward:                       v.LevelReward,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}

func (u *UserRepo) GetUserV1BoundLast(ctx context.Context) (*biz.UserV1Bound, error) {
	var v UserV1Bound

	if err := u.data.DB(ctx).Table("user_v1_bound_event").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "USER_V1_BOUND_ERROR", err.Error())
	}

	return toBizUserV1Bound(&v), nil
}

func (u *UserRepo) GetUserV1BoundByAddress(ctx context.Context, address string) (*biz.UserV1Bound, error) {
	var v UserV1Bound

	if err := u.data.DB(ctx).Table("user_v1_bound_event").Where("user_addr = ?", address).First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "USER_V1_BOUND_BY_ADDRESS_ERROR", err.Error())
	}

	return toBizUserV1Bound(&v), nil
}

func (u *UserRepo) GetUserV1Bounds(ctx context.Context) ([]*biz.UserV1Bound, error) {
	var rows []UserV1Bound
	if err := u.data.DB(ctx).Table("user_v1_bound_event").Order("id asc").Find(&rows).Error; nil != err {
		return nil, errors.New(500, "USER_V1_BOUNDS_ERROR", err.Error())
	}

	result := make([]*biz.UserV1Bound, 0, len(rows))
	for i := range rows {
		result = append(result, toBizUserV1Bound(&rows[i]))
	}
	return result, nil
}

func (u *UserRepo) GetUserV1BoundsByIDs(ctx context.Context, ids []uint64) ([]*biz.UserV1Bound, error) {
	if 0 == len(ids) {
		return []*biz.UserV1Bound{}, nil
	}

	var rows []UserV1Bound
	if err := u.data.DB(ctx).Table("user_v1_bound_event").Where("id IN ?", ids).Order("id asc").Find(&rows).Error; nil != err {
		return nil, errors.New(500, "USER_V1_BOUNDS_BY_IDS_ERROR", err.Error())
	}

	result := make([]*biz.UserV1Bound, 0, len(rows))
	for i := range rows {
		result = append(result, toBizUserV1Bound(&rows[i]))
	}
	return result, nil
}

func decimalOrZero(value string) string {
	if "" == strings.TrimSpace(value) {
		return "0"
	}
	return value
}

func (u *UserRepo) InsertUserV1Bound(ctx context.Context, iData *biz.UserV1Bound) error {
	var s UserV1Bound

	s.ID = iData.ID
	s.BlockNumber = iData.BlockNumber

	s.UserAddr = iData.UserAddr
	s.ParentAddr = iData.ParentAddr
	s.RecommendCode = iData.RecommendCode
	s.Amount = decimalOrZero(iData.Amount)
	s.AmountHistory = decimalOrZero(iData.AmountHistory)
	s.ChildrenAmount = decimalOrZero(iData.ChildrenAmount)
	s.ChildrenAmountHistory = decimalOrZero(iData.ChildrenAmountHistory)
	s.ChildrenAmountExtra = decimalOrZero(iData.ChildrenAmountExtra)
	s.RewardRecommendAmount = decimalOrZero(iData.RewardRecommendAmount)
	s.RewardRecommendPay = decimalOrZero(iData.RewardRecommendPay)
	s.RewardRecommendStoreAmount = decimalOrZero(iData.RewardRecommendStoreAmount)
	s.RewardRecommendFee = decimalOrZero(iData.RewardRecommendFee)
	s.RewardRecommendTeamUAmount = decimalOrZero(iData.RewardRecommendTeamUAmount)
	s.RewardRecommendClaimedTeamUNet = decimalOrZero(iData.RewardRecommendClaimedTeamUNet)
	s.RewardRecommendClaimedTeamUAmount = decimalOrZero(iData.RewardRecommendClaimedTeamUAmount)
	s.RewardRecommendClaimedTeamUFee = decimalOrZero(iData.RewardRecommendClaimedTeamUFee)
	s.RewardRecommendExpired = decimalOrZero(iData.RewardRecommendExpired)
	s.LineU = decimalOrZero(iData.LineU)
	s.LineCoinU = decimalOrZero(iData.LineCoinU)
	s.LineCoin = decimalOrZero(iData.LineCoin)
	s.LineFee = decimalOrZero(iData.LineFee)
	s.LevelReward = decimalOrZero(iData.LevelReward)

	if err := u.data.DB(ctx).Table("user_v1_bound_event").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_USER_V1_BOUND_ERROR", "信息创建失败")
	}
	iData.ID = s.ID

	return nil
}

func (u *UserRepo) DeleteUserV1BoundAll(ctx context.Context) error {
	if err := u.data.DB(ctx).Exec("DELETE FROM user_v1_bound_event").Error; err != nil {
		return errors.New(500, "DELETE_USER_V1_BOUND_ERROR", "历史数据清理失败")
	}
	return nil
}

func (u *UserRepo) GetUserV1BoundSyncProgress(ctx context.Context) (*biz.UserV1BoundSyncProgress, error) {
	var v UserV1BoundSyncProgress

	if err := u.data.DB(ctx).Table("user_v1_bound_sync_progress").Where("id = ?", 1).First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "USER_V1_BOUND_SYNC_PROGRESS_ERROR", err.Error())
	}

	return &biz.UserV1BoundSyncProgress{
		ID:                 v.ID,
		LastProcessedBlock: v.LastProcessedBlock,
		CreatedAt:          v.CreatedAt,
		UpdatedAt:          v.UpdatedAt,
	}, nil
}

func (u *UserRepo) SaveUserV1BoundSyncProgress(ctx context.Context, lastProcessedBlock uint64) error {
	progress := UserV1BoundSyncProgress{
		ID:                 1,
		LastProcessedBlock: lastProcessedBlock,
	}

	if err := u.data.DB(ctx).Table("user_v1_bound_sync_progress").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"last_processed_block", "updated_at"}),
	}).Create(&progress).Error; err != nil {
		return errors.New(500, "SAVE_USER_V1_BOUND_SYNC_PROGRESS_ERROR", "同步进度保存失败")
	}

	return nil
}

func (u *UserRepo) GetUserV1PerformanceSyncProgress(ctx context.Context, streamName string) (*biz.UserV1PerformanceSyncProgress, error) {
	var v UserV1PerformanceSyncProgress
	if err := u.data.DB(ctx).Table("user_v1_performance_sync_progress").Where("stream_name = ?", streamName).First(&v).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "USER_V1_PERFORMANCE_SYNC_PROGRESS_ERROR", err.Error())
	}

	return &biz.UserV1PerformanceSyncProgress{
		StreamName:         v.StreamName,
		LastProcessedBlock: v.LastProcessedBlock,
		CreatedAt:          v.CreatedAt,
		UpdatedAt:          v.UpdatedAt,
	}, nil
}

func (u *UserRepo) SaveUserV1PerformanceSyncProgress(ctx context.Context, streamName string, lastProcessedBlock uint64) error {
	progress := UserV1PerformanceSyncProgress{
		StreamName:         streamName,
		LastProcessedBlock: lastProcessedBlock,
	}
	if err := u.data.DB(ctx).Table("user_v1_performance_sync_progress").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "stream_name"}},
		DoUpdates: clause.AssignmentColumns([]string{"last_processed_block", "updated_at"}),
	}).Create(&progress).Error; nil != err {
		return errors.New(500, "SAVE_USER_V1_PERFORMANCE_SYNC_PROGRESS_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) DeleteUserV1PerformanceSyncProgress(ctx context.Context, streamName string) error {
	if err := u.data.DB(ctx).Table("user_v1_performance_sync_progress").Where("stream_name = ?", streamName).Delete(&UserV1PerformanceSyncProgress{}).Error; nil != err {
		return errors.New(500, "DELETE_USER_V1_PERFORMANCE_SYNC_PROGRESS_ERROR", err.Error())
	}
	return nil
}

func createPerformanceEvent(ctx context.Context, db *gorm.DB, table string, value interface{}) (bool, error) {
	result := db.WithContext(ctx).Table(table).Clauses(clause.OnConflict{DoNothing: true}).Create(value)
	if nil != result.Error {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (u *UserRepo) InsertUserV1StakeChanged(ctx context.Context, event *biz.UserV1StakeChanged) (bool, error) {
	row := &UserV1StakeChanged{
		BlockNumber: event.BlockNumber,
		EventKey:    event.EventKey,
		TxHash:      event.TxHash,
		UserAddr:    event.UserAddr,
		Amount:      event.Amount,
		IsAdd:       event.IsAdd,
	}
	inserted, err := createPerformanceEvent(ctx, u.data.DB(ctx), "user_v1_stake_changed_event", row)
	if nil != err {
		return false, errors.New(500, "CREATE_USER_V1_STAKE_CHANGED_ERROR", err.Error())
	}
	event.ID = row.ID
	return inserted, nil
}

func (u *UserRepo) InsertUserV1ExtraChanged(ctx context.Context, event *biz.UserV1ExtraChanged) (bool, error) {
	row := &UserV1ExtraChanged{
		BlockNumber: event.BlockNumber,
		EventKey:    event.EventKey,
		TxHash:      event.TxHash,
		UserAddr:    event.UserAddr,
		ExtraAmount: event.ExtraAmount,
	}
	inserted, err := createPerformanceEvent(ctx, u.data.DB(ctx), "user_v1_extra_changed_event", row)
	if nil != err {
		return false, errors.New(500, "CREATE_USER_V1_EXTRA_CHANGED_ERROR", err.Error())
	}
	event.ID = row.ID
	return inserted, nil
}

func (u *UserRepo) InsertStakingV1TeamBooked(ctx context.Context, event *biz.StakingV1Reward) (bool, error) {
	row := &StakingV1TeamBooked{
		BlockNumber: event.BlockNumber,
		EventKey:    event.EventKey,
		TxHash:      event.TxHash,
		FromAddr:    event.FromAddr,
		ToAddr:      event.ToAddr,
		Amount:      event.Amount,
		StoreAmount: event.StoreAmount,
		Pay:         event.Pay,
		Fee:         event.Fee,
	}
	inserted, err := createPerformanceEvent(ctx, u.data.DB(ctx), "staking_v1_team_booked_event", row)
	if nil != err {
		return false, errors.New(500, "CREATE_STAKING_V1_TEAM_BOOKED_ERROR", err.Error())
	}
	event.ID = row.ID
	return inserted, nil
}

func (u *UserRepo) InsertStakingV1TeamClaimed(ctx context.Context, event *biz.StakingV1Reward) (bool, error) {
	row := &StakingV1TeamClaimed{
		BlockNumber: event.BlockNumber,
		EventKey:    event.EventKey,
		TxHash:      event.TxHash,
		UserAddr:    event.UserAddr,
		Amount:      event.Amount,
		Fee:         event.Fee,
		Net:         event.Net,
	}
	inserted, err := createPerformanceEvent(ctx, u.data.DB(ctx), "staking_v1_team_claimed_event", row)
	if nil != err {
		return false, errors.New(500, "CREATE_STAKING_V1_TEAM_CLAIMED_ERROR", err.Error())
	}
	event.ID = row.ID
	return inserted, nil
}

func (u *UserRepo) InsertStakingV1TeamExpired(ctx context.Context, event *biz.StakingV1Reward) (bool, error) {
	row := &StakingV1TeamExpired{
		BlockNumber: event.BlockNumber,
		EventKey:    event.EventKey,
		TxHash:      event.TxHash,
		FromAddr:    event.FromAddr,
		ToAddr:      event.ToAddr,
		Amount:      event.Amount,
	}
	inserted, err := createPerformanceEvent(ctx, u.data.DB(ctx), "staking_v1_team_expired_event", row)
	if nil != err {
		return false, errors.New(500, "CREATE_STAKING_V1_TEAM_EXPIRED_ERROR", err.Error())
	}
	event.ID = row.ID
	return inserted, nil
}

func (u *UserRepo) InsertStakingV1LineClaimed(ctx context.Context, event *biz.StakingV1Reward) (bool, error) {
	row := &StakingV1LineClaimed{
		BlockNumber: event.BlockNumber,
		EventKey:    event.EventKey,
		TxHash:      event.TxHash,
		UserAddr:    event.UserAddr,
		OrderID:     event.OrderID,
		GrossU:      event.GrossU,
		FeeU:        event.FeeU,
		PaidMs:      event.PaidMs,
		MsAmount:    event.MsAmount,
	}
	inserted, err := createPerformanceEvent(ctx, u.data.DB(ctx), "staking_v1_line_claimed_event", row)
	if nil != err {
		return false, errors.New(500, "CREATE_STAKING_V1_LINE_CLAIMED_ERROR", err.Error())
	}
	event.ID = row.ID
	return inserted, nil
}

func decimalAdd(column string, amount string) clause.Expr {
	return gorm.Expr(column+" + CAST(? AS DECIMAL(65,18))", amount)
}

func decimalSub(column string, amount string) clause.Expr {
	return gorm.Expr(column+" - CAST(? AS DECIMAL(65,18))", amount)
}

func (u *UserRepo) UpdateUserV1StakeAmount(ctx context.Context, userID uint64, amount string, isAdd bool) error {
	if isAdd {
		updates := map[string]interface{}{
			"amount":         decimalAdd("amount", amount),
			"amount_history": decimalAdd("amount_history", amount),
		}
		if err := u.data.DB(ctx).Table("user_v1_bound_event").Where("id = ?", userID).Updates(updates).Error; nil != err {
			return errors.New(500, "UPDATE_USER_V1_STAKE_AMOUNT_ERROR", err.Error())
		}
		return nil
	}

	result := u.data.DB(ctx).Table("user_v1_bound_event").
		Where("id = ? AND amount >= CAST(? AS DECIMAL(65,18))", userID, amount).
		Update("amount", decimalSub("amount", amount))
	if nil != result.Error {
		return errors.New(500, "UPDATE_USER_V1_STAKE_AMOUNT_ERROR", result.Error.Error())
	}
	if 1 != result.RowsAffected {
		err := fmt.Errorf("user %d amount is less than unstake amount %s", userID, amount)
		return errors.New(500, "UPDATE_USER_V1_STAKE_AMOUNT_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) UpdateUserV1ChildrenAmount(ctx context.Context, userIDs []uint64, amount string, isAdd bool) error {
	if 0 == len(userIDs) {
		return nil
	}
	if isAdd {
		updates := map[string]interface{}{
			"children_amount":         decimalAdd("children_amount", amount),
			"children_amount_history": decimalAdd("children_amount_history", amount),
		}
		if err := u.data.DB(ctx).Table("user_v1_bound_event").Where("id IN ?", userIDs).Updates(updates).Error; nil != err {
			return errors.New(500, "UPDATE_USER_V1_CHILDREN_AMOUNT_ERROR", err.Error())
		}
		return nil
	}

	result := u.data.DB(ctx).Table("user_v1_bound_event").
		Where("id IN ? AND children_amount >= CAST(? AS DECIMAL(65,18))", userIDs, amount).
		Update("children_amount", decimalSub("children_amount", amount))
	if nil != result.Error {
		return errors.New(500, "UPDATE_USER_V1_CHILDREN_AMOUNT_ERROR", result.Error.Error())
	}
	if int64(len(userIDs)) != result.RowsAffected {
		err := fmt.Errorf("one or more ancestors have children_amount less than unstake amount %s", amount)
		return errors.New(500, "UPDATE_USER_V1_CHILDREN_AMOUNT_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) UpdateUserV1ExtraAmount(ctx context.Context, userID uint64, amount string) error {
	if err := u.data.DB(ctx).Table("user_v1_bound_event").Where("id = ?", userID).Update("children_amount_extra", amount).Error; nil != err {
		return errors.New(500, "UPDATE_USER_V1_EXTRA_AMOUNT_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) UpdateUserV1TeamBooked(ctx context.Context, userID uint64, event *biz.StakingV1Reward) error {
	updates := map[string]interface{}{
		"reward_recommend_amount":        decimalAdd("reward_recommend_amount", event.Amount),
		"reward_recommend_pay":           decimalAdd("reward_recommend_pay", event.Pay),
		"reward_recommend_store_amount":  decimalAdd("reward_recommend_store_amount", event.StoreAmount),
		"reward_recommend_fee":           decimalAdd("reward_recommend_fee", event.Fee),
		"reward_recommend_team_u_amount": decimalAdd("reward_recommend_team_u_amount", event.StoreAmount),
	}
	if err := u.data.DB(ctx).Table("user_v1_bound_event").Where("id = ?", userID).Updates(updates).Error; nil != err {
		return errors.New(500, "UPDATE_USER_V1_TEAM_BOOKED_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) UpdateUserV1TeamClaimed(ctx context.Context, userID uint64, event *biz.StakingV1Reward) error {
	updates := map[string]interface{}{
		"reward_recommend_claimed_team_u_net":    decimalAdd("reward_recommend_claimed_team_u_net", event.Net),
		"reward_recommend_claimed_team_u_amount": decimalAdd("reward_recommend_claimed_team_u_amount", event.Amount),
		"reward_recommend_claimed_team_u_fee":    decimalAdd("reward_recommend_claimed_team_u_fee", event.Fee),
		"reward_recommend_team_u_amount":         gorm.Expr("reward_recommend_team_u_amount - CAST(? AS DECIMAL(65,18))", event.Amount),
	}
	result := u.data.DB(ctx).Table("user_v1_bound_event").
		Where("id = ? AND reward_recommend_team_u_amount >= CAST(? AS DECIMAL(65,18))", userID, event.Amount).
		Updates(updates)
	if nil != result.Error {
		return errors.New(500, "UPDATE_USER_V1_TEAM_CLAIMED_ERROR", result.Error.Error())
	}
	if 1 != result.RowsAffected {
		err := fmt.Errorf("user %d team_u is less than TeamClaimed amount %s", userID, event.Amount)
		return errors.New(500, "UPDATE_USER_V1_TEAM_CLAIMED_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) UpdateUserV1TeamExpired(ctx context.Context, userID uint64, amount string) error {
	if err := u.data.DB(ctx).Table("user_v1_bound_event").Where("id = ?", userID).Update("reward_recommend_expired", decimalAdd("reward_recommend_expired", amount)).Error; nil != err {
		return errors.New(500, "UPDATE_USER_V1_TEAM_EXPIRED_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) UpdateUserV1LineClaimed(ctx context.Context, userID uint64, event *biz.StakingV1Reward) error {
	updates := map[string]interface{}{
		"line_fee": decimalAdd("line_fee", event.FeeU),
	}
	if event.PaidMs {
		updates["line_coin_u"] = decimalAdd("line_coin_u", event.GrossU)
		updates["line_coin"] = decimalAdd("line_coin", event.MsAmount)
	} else {
		updates["line_u"] = decimalAdd("line_u", event.GrossU)
	}
	if err := u.data.DB(ctx).Table("user_v1_bound_event").Where("id = ?", userID).Updates(updates).Error; nil != err {
		return errors.New(500, "UPDATE_USER_V1_LINE_CLAIMED_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) UpdateUserV1LevelReward(ctx context.Context, userID uint64, amount string) error {
	if err := u.data.DB(ctx).Table("user_v1_bound_event").Where("id = ?", userID).Update("level_reward", amount).Error; nil != err {
		return errors.New(500, "UPDATE_USER_V1_LEVEL_REWARD_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) ResetUserV1Performance(ctx context.Context) error {
	columns := []string{
		"amount", "amount_history", "children_amount", "children_amount_history", "children_amount_extra",
		"reward_recommend_amount", "reward_recommend_pay", "reward_recommend_store_amount", "reward_recommend_fee",
		"reward_recommend_team_u_amount", "reward_recommend_claimed_team_u_net", "reward_recommend_claimed_team_u_amount",
		"reward_recommend_claimed_team_u_fee", "reward_recommend_expired", "line_u", "line_coin_u", "line_coin",
		"line_fee", "level_reward",
	}
	updates := make(map[string]interface{}, len(columns))
	for _, column := range columns {
		updates[column] = 0
	}
	if err := u.data.DB(ctx).Table("user_v1_bound_event").Where("1 = 1").Updates(updates).Error; nil != err {
		return errors.New(500, "RESET_USER_V1_PERFORMANCE_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) DeleteUserV1PerformanceEvents(ctx context.Context) error {
	tables := []string{
		"user_v1_stake_changed_event", "user_v1_extra_changed_event", "staking_v1_team_booked_event",
		"staking_v1_team_claimed_event", "staking_v1_team_expired_event", "staking_v1_line_claimed_event",
	}
	for _, table := range tables {
		if err := u.data.DB(ctx).Exec("DELETE FROM " + table).Error; nil != err {
			return errors.New(500, "DELETE_USER_V1_PERFORMANCE_EVENTS_ERROR", err.Error())
		}
	}
	return nil
}

func (u *UserRepo) GetStakingStakedLast(ctx context.Context) (*biz.StakingStaked, error) {
	var v StakingStaked

	if err := u.data.DB(ctx).Table("staking_staked_event").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "STAKING_STAKED_ERROR", err.Error())
	}

	return &biz.StakingStaked{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		UserAddr:   v.UserAddr,
		Amount:     v.Amount,
		Timestamp:  v.Timestamp,
		StakeIndex: v.StakeIndex,
		Duration:   v.Duration,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,

		CheckStatus: v.CheckStatus,
		CheckTime:   v.CheckTime,
	}, nil
}

func (u *UserRepo) InsertStakingStaked(ctx context.Context, iData *biz.StakingStaked) error {
	var s StakingStaked

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.UserAddr = iData.UserAddr
	s.Amount = iData.Amount
	s.Timestamp = iData.Timestamp
	s.StakeIndex = iData.StakeIndex
	s.Duration = iData.Duration

	if err := u.data.DB(ctx).Table("staking_staked_event").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_STAKING_STAKED_ERROR", "信息创建失败")
	}

	return nil
}
func (u *UserRepo) GetStakingUnstakedLast(ctx context.Context) (*biz.StakingUnstaked, error) {
	var v StakingUnstaked

	if err := u.data.DB(ctx).Table("staking_unstaked_event").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "STAKING_UNSTAKED_ERROR", err.Error())
	}

	return &biz.StakingUnstaked{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		UserAddr:   v.UserAddr,
		Amount:     v.Amount,
		Timestamp:  v.Timestamp,
		StakeIndex: v.StakeIndex,
		Reward:     v.Reward,
		TTL:        v.TTL,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,

		CheckStatus: v.CheckStatus,
		CheckTime:   v.CheckTime,
	}, nil
}

func (u *UserRepo) InsertStakingUnstaked(ctx context.Context, iData *biz.StakingUnstaked) error {
	var s StakingUnstaked

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.UserAddr = iData.UserAddr
	s.Amount = iData.Amount
	s.Timestamp = iData.Timestamp
	s.StakeIndex = iData.StakeIndex
	s.Reward = iData.Reward
	s.TTL = iData.TTL

	if err := u.data.DB(ctx).Table("staking_unstaked_event").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_STAKING_UNSTAKED_ERROR", "信息创建失败")
	}

	return nil
}

func (u *UserRepo) GetStakingQueueAddedLast(ctx context.Context) (*biz.StakingQueueAdded, error) {
	var v StakingQueueAdded

	if err := u.data.DB(ctx).Table("staking_queue_added_event").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "STAKING_QUEUE_ADDED_ERROR", err.Error())
	}

	return &biz.StakingQueueAdded{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		QueueIndex: v.QueueIndex,
		UserAddr:   v.UserAddr,
		Amount:     v.Amount,
		StakeIndex: v.StakeIndex,
		QueuedAt:   v.QueuedAt,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,

		CheckStatus: v.CheckStatus,
		CheckTime:   v.CheckTime,
	}, nil
}

func (u *UserRepo) InsertStakingQueueAdded(ctx context.Context, iData *biz.StakingQueueAdded) error {
	var s StakingQueueAdded

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.QueueIndex = iData.QueueIndex
	s.UserAddr = iData.UserAddr
	s.Amount = iData.Amount
	s.StakeIndex = iData.StakeIndex
	s.QueuedAt = iData.QueuedAt

	if err := u.data.DB(ctx).Table("staking_queue_added_event").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_STAKING_QUEUE_ADDED_ERROR", "信息创建失败")
	}

	return nil
}
