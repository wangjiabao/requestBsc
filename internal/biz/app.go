package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"math/big"
	pb "requestEth/api/requestEth/v1"
	"strconv"
	"strings"
	"time"
)

type SwapTrade struct {
	ID              uint64
	BlockNumber     uint64
	LogIndex        uint32
	BlockTime       uint64
	Sender          string
	ToAddr          string
	Side            uint8
	Amount0In       float64
	Amount1In       float64
	Amount0OutNet   float64
	Amount1OutGross float64
	Amount0OutGross float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type PrimaryBuy struct {
	ID           uint64
	BlockNumber  uint64
	BlockTime    uint64
	LogIndex     uint32
	Buyer        string
	ToAddr       string
	UsdtUsed     float64
	AusdGrossOut float64
	AusdFee      float64
	AusdNetOut   float64
	PriceBefore  float64
	PriceAfter   float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type PrimarySell struct {
	ID          uint64
	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint32
	Seller      string
	ToAddr      string
	AusdGrossIn float64
	AusdFee     float64
	AusdBurn    float64
	UsdtOut     float64
	PriceBefore float64
	PriceAfter  float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RewardNotified struct {
	ID               uint64
	BlockNumber      uint64
	BlockTime        uint64
	LogIndex         uint32
	User             string
	L1               string
	L2               string
	Profit           float64
	UserShare        float64
	Top              string
	Pool             float64
	UplinePortionBps uint64
	ToL1             float64
	ToL2             float64
	ToTop            float64
	ToProject        float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type NftMarketPurchase struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	Buyer   string
	Seller  string
	TokenID uint64

	PriceUSDT  float64
	FeePaidInB uint8
	FeeUSDT    float64
	FeeB       float64

	CreatedAt   time.Time
	UpdatedAt   time.Time
	CheckStatus uint64
}

type RewardDetail struct {
	ID         uint64
	User       string
	Amount     float64
	Reason     uint64
	NotifiedId uint64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	BlockTime  uint64
}

type NftMinted struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	ToAddr   string
	MintAddr string
	TokenID  uint64

	Tier     uint64
	UsdtPaid float64

	Status   uint8
	ListedAt uint64

	OpenStatus uint8
	OpenedAt   uint64

	CreatedAt   time.Time
	UpdatedAt   time.Time
	CheckStatus uint64
	CheckTime   uint64
}

type NftMarketListed struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	Seller    string
	TokenID   uint64
	Timestamp uint64

	CreatedAt   time.Time
	UpdatedAt   time.Time
	CheckStatus uint64
}

type NftMarketUnlisted struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	Operator string
	TokenID  uint64

	CreatedAt time.Time
	UpdatedAt time.Time

	CheckStatus uint64
}

type NftOpened struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	UserAddr string
	TokenID  uint64

	OpenedAt uint64
	Reward   float64

	CreatedAt   time.Time
	UpdatedAt   time.Time
	CheckStatus uint64
}

type NftTransfer struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	FromAddr string
	ToAddr   string
	TokenID  uint64

	CreatedAt   time.Time
	UpdatedAt   time.Time
	CheckStatus uint64
}

type UserRegistered struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	UserAddr   string
	ParentAddr string
	TopAddr    string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type BindReferral struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	UserAddr   string
	ParentAddr string

	CreatedAt time.Time
	UpdatedAt time.Time

	CheckStatus uint64
	CheckTime   uint64
	Level       int8
}

type UserV1Bound struct {
	ID uint64

	BlockNumber uint64

	UserAddr      string
	Name          string
	ParentAddr    string
	RecommendCode string

	Amount                            string
	AmountHistory                     string
	InvestmentCount                   uint64
	ChildrenAmount                    string
	ChildrenAmountHistory             string
	ChildrenAmountExtra               string
	RewardRecommendAmount             string
	RewardRecommendPay                string
	RewardRecommendStoreAmount        string
	RewardRecommendFee                string
	RewardRecommendTeamUAmount        string
	RewardRecommendClaimedTeamUNet    string
	RewardRecommendClaimedTeamUAmount string
	RewardRecommendClaimedTeamUFee    string
	RewardRecommendExpired            string
	LineU                             string
	LineCoinU                         string
	LineCoin                          string
	LineFee                           string
	LevelReward                       string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserV1BoundSyncProgress struct {
	ID                 uint8
	LastProcessedBlock uint64
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type UserV1PerformanceSyncProgress struct {
	StreamName         string
	LastProcessedBlock uint64
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

const (
	UserV1PerformanceStreamStake         = "stake_changed"
	UserV1PerformanceStreamExtra         = "extra_changed"
	UserV1PerformanceStreamReward        = "staking_reward"
	UserV1PerformanceStreamOrder         = "staking_order"
	UserV1PerformanceStreamOrderRecovery = "staking_order_recovery"
	UserV1PerformanceStreamOrderTarget   = "staking_order_recovery_target"
	UserV1PerformanceStreamRecovery      = "recovery"
)

type UserV1StakeChanged struct {
	ID               uint64
	BlockNumber      uint64
	BlockTime        uint64
	LogIndex         uint
	EventKey         string
	TxHash           string
	UserAddr         string
	Amount           string
	IsAdd            bool
	InvestmentNumber uint64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type UserV1Overview struct {
	RegisteredUserCount                  uint64
	HistoricalInvestorCount              uint64
	CurrentInvestorCount                 uint64
	CurrentAmountGTE10000UserCount       uint64
	HistoricalAmountGTE10000UserCount    uint64
	InvestmentCountGT2UserCount          uint64
	TodayInvestmentAmount                string
	TodayInvestmentOrderCount            uint64
	YesterdayInvestmentAmount            string
	YesterdayInvestmentOrderCount        uint64
	TodayReinvestmentAmount              string
	MissingInvestmentBlockTimeEventCount uint64
	MissingInvestmentNumberEventCount    uint64
}

type UserV1ExtraChanged struct {
	ID          uint64
	BlockNumber uint64
	LogIndex    uint
	EventKey    string
	TxHash      string
	UserAddr    string
	ExtraAmount string
	ApplyStatus uint8
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

const (
	UserV1ExtraChangedApplyStatusApplied      uint8 = 1
	UserV1ExtraChangedApplyStatusUnregistered uint8 = 2
)

const (
	StakingV1RewardTeamBooked  = "team_booked"
	StakingV1RewardTeamClaimed = "team_claimed"
	StakingV1RewardTeamExpired = "team_expired"
	StakingV1RewardLineClaimed = "line_claimed"
)

type StakingV1Reward struct {
	ID          uint64
	BlockNumber uint64
	LogIndex    uint
	EventKey    string
	TxHash      string
	EventType   string

	FromAddr string
	ToAddr   string
	UserAddr string

	Amount      string
	StoreAmount string
	Pay         string
	Fee         string
	Net         string

	OrderID  string
	GrossU   string
	FeeU     string
	PaidMs   bool
	MsAmount string

	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	StakingV1OrderStatusQueued  uint8 = 1
	StakingV1OrderStatusRunning uint8 = 2
	StakingV1OrderStatusExited  uint8 = 3

	StakingV1OrderEventCreated   = "created"
	StakingV1OrderEventEntered   = "entered"
	StakingV1OrderEventExited    = "exited"
	StakingV1OrderEventCapSet    = "cap_set"
	StakingV1OrderEventQueued    = "queued"
	StakingV1OrderEventQueueDone = "queue_done"
	StakingV1OrderEventPlanSet   = "plan_set"
)

// StakingV1Order is the latest locally known state of one staking order.
// uint256 identifiers and token amounts stay as decimal strings so no chain
// precision is lost before they are written to DECIMAL(65,0/18) columns.
type StakingV1Order struct {
	ID             uint64
	OrderID        string
	UserID         uint64
	UserAddr       string
	UserOrderIndex string

	Amount        string
	BaseCap       string
	Cap           string
	Used          string
	Remaining     string
	Compensation  string
	LinePaid      string
	LineClaimable string
	PlanID        string

	CreatedTime    uint64
	StartTime      uint64
	ClaimEffective uint64
	DaysCount      uint32
	Status         uint8

	QueueIndex string
	QueueLiqU  string
	QueuedAt   uint64
	QueueDone  bool

	CreatedBlock    uint64
	EnteredBlock    uint64
	ExitedBlock     uint64
	LastSyncedBlock uint64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// StakingV1OrderEvent is the common in-memory representation of the order and
// plan lifecycle events. Only the fields belonging to EventType are persisted.
type StakingV1OrderEvent struct {
	ID          uint64
	BlockNumber uint64
	LogIndex    uint
	EventKey    string
	TxHash      string
	EventType   string
	OrderID     string
	UserAddr    string
	UserID      uint64

	UserOrderIndex string
	Amount         string
	Cap            string
	PlanID         string
	MinAmount      string
	MaxAmount      string
	OutAmount      string
	DaysCount      uint32
	Enabled        bool
	StartTime      uint64
	Used           string
	OldCap         string
	NewCap         string
	QueueIndex     string
	QueueLiqU      string
	QueuedAt       uint64

	CreatedAt time.Time
	UpdatedAt time.Time
}

// StakingV1OrderSnapshot contains an absolute state read from the contract.
// Applying a snapshot is monotonic by LastSyncedBlock.
type StakingV1OrderSnapshot struct {
	OrderID        string
	UserID         uint64
	UserAddr       string
	UserOrderIndex string

	Amount        string
	BaseCap       string
	Cap           string
	Used          string
	Remaining     string
	Compensation  string
	LinePaid      string
	LineClaimable string
	PlanID        string

	CreatedTime    uint64
	StartTime      uint64
	ClaimEffective uint64
	DaysCount      uint32
	Status         uint8

	QueueIndex string
	QueueLiqU  string
	QueuedAt   uint64
	QueueDone  bool

	CreatedBlock    uint64
	EnteredBlock    uint64
	ExitedBlock     uint64
	LastSyncedBlock uint64
}

type StakingV1OrderQuery struct {
	Page     uint64
	PageSize uint64
	UserID   uint64
	Address  string
	Status   uint8
	OrderBy  string
	Order    string
}

type StakingV1OrderUser struct {
	UserID   uint64
	UserAddr string
}

type StakingV1OrderIntegrity struct {
	MasterWithoutCreated      uint64
	CreatedWithoutMaster      uint64
	ExitedWithoutExit         uint64
	ExitNotMarkedExited       uint64
	RunningWithoutEntered     uint64
	QueuedWithEntered         uint64
	QueueDoneWithoutQueued    uint64
	QueueDoneWithoutEntered   uint64
	LifecycleIdentityMismatch uint64
	MasterCreatedMismatch     uint64
	DuplicateCreatedOrderID   uint64
	DuplicateExitOrderID      uint64
	CreatedCount              uint64
	MinCreatedOrderID         string
	MaxCreatedOrderID         string
}

type StakingStaked struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	UserAddr   string
	Amount     float64
	Timestamp  uint64
	StakeIndex uint64
	Duration   uint64

	CreatedAt time.Time
	UpdatedAt time.Time

	CheckStatus uint64
	CheckTime   uint64
}

type StakingUnstaked struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	UserAddr   string
	Amount     float64
	Timestamp  uint64
	StakeIndex uint64
	Reward     float64
	TTL        uint64

	CreatedAt time.Time
	UpdatedAt time.Time

	CheckStatus uint64
	CheckTime   uint64
}

type StakingQueueAdded struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	QueueIndex uint64
	UserAddr   string
	Amount     float64
	StakeIndex uint8
	QueuedAt   uint64

	CreatedAt time.Time
	UpdatedAt time.Time

	CheckStatus uint64
	CheckTime   uint64
}

type UserRepo interface {
	GetSwapTradeLast(ctx context.Context) (*SwapTrade, error)
	GetSwapTrade(ctx context.Context, start, end uint64) ([]*SwapTrade, error)
	InsertSwapTrade(ctx context.Context, iData *SwapTrade) error
	GetPrimaryBuyLast(ctx context.Context) (*PrimaryBuy, error)
	GetPrimaryBuy(ctx context.Context, start, end uint64) ([]*PrimaryBuy, error)
	InsertPrimaryBuy(ctx context.Context, iData *PrimaryBuy) error
	GetPrimarySellLast(ctx context.Context) (*PrimarySell, error)
	GetPrimarySell(ctx context.Context, start, end uint64) ([]*PrimarySell, error)
	InsertPrimarySell(ctx context.Context, iData *PrimarySell) error
	GetRewardNotifiedLast(ctx context.Context) (*RewardNotified, error)
	GetRewardNotified(ctx context.Context, start, end uint64) ([]*RewardNotified, error)
	GetRewardNotifiedByIds(ctx context.Context, ids []uint64) (map[uint64]*RewardNotified, error)
	InsertRewardNotified(ctx context.Context, iData *RewardNotified) error
	GetUserRewardByUserIdPage(ctx context.Context, b *Pagination, address string, reason uint64) ([]*RewardDetail, error, int64)
	GetNftMarketPurchaseByAddressPage(ctx context.Context, b *Pagination, address string, addressTwo []string, side uint64) ([]*NftMarketPurchase, error, int64)
	GetNftMarketPurchaseLast(ctx context.Context) (*NftMarketPurchase, error)
	InsertNftMarketPurchase(ctx context.Context, iData *NftMarketPurchase) error
	InsertNftMinted(ctx context.Context, iData *NftMinted) error
	UpdateNftMinted(ctx context.Context, tokenId uint64, mintAddr string) error
	GetNftMintedLast(ctx context.Context) (*NftMinted, error)
	InsertNftMarketListed(ctx context.Context, iData *NftMarketListed) error
	GetNftMarketListedLast(ctx context.Context) (*NftMarketListed, error)
	InsertNftMarketUnlisted(ctx context.Context, iData *NftMarketUnlisted) error
	GetNftMarketUnlistedLast(ctx context.Context) (*NftMarketUnlisted, error)
	InsertNftOpened(ctx context.Context, iData *NftOpened) error
	GetNftOpenedLast(ctx context.Context) (*NftOpened, error)
	InsertNftTransfer(ctx context.Context, iData *NftTransfer) error
	GetNftTransferLast(ctx context.Context) (*NftTransfer, error)
	GetNftTransferLastNoCheck(ctx context.Context) ([]*NftTransfer, error)
	GetNftMintedByTokenIds(ctx context.Context, tokenIds []uint64) (map[uint64]*NftMinted, error)
	UpdateNftMintedToAddress(ctx context.Context, id, idT, check uint64, toAddr string) error
	GetNftListLastNoCheck(ctx context.Context) ([]*NftMarketListed, error)
	GetNftUnListLastNoCheck(ctx context.Context) ([]*NftMarketUnlisted, error)
	GetNftBuyLastNoCheck(ctx context.Context) ([]*NftMarketPurchase, error)
	GetNftOpenLastNoCheck(ctx context.Context) ([]*NftOpened, error)
	UpdateNftMintedListStatus(ctx context.Context, id, idT, checkTime uint64) error
	UpdateNftMintedUnListStatus(ctx context.Context, id, idT, checkTime uint64) error
	UpdateNftMintedBuyStatus(ctx context.Context, id, idT, checkTime uint64) error
	UpdateNftMintedOpenStatus(ctx context.Context, id, idT, checkTime uint64) error
	UpdateUnlistedCheckStatus(ctx context.Context, id, idT, checkTime uint64) error
	UpdateListedCheckStatus(ctx context.Context, id, idT, checkTime uint64) error
	UpdateBuyCheckStatus(ctx context.Context, id, idT, checkTime uint64) error
	GetNftMintedByAddressPage(ctx context.Context, b *Pagination, start uint64, end uint64, address []string, status uint64, order uint64, tier uint64, openStatus uint64, openAtOrder uint64) ([]*NftMinted, error, int64)
	GetNftMintedPage(ctx context.Context, b *Pagination, order uint64, orderTwo uint64, tier uint64) ([]*NftMinted, error, int64)
	GetUserRegisteredLast(ctx context.Context) (*UserRegistered, error)
	InsertUserRegistered(ctx context.Context, iData *UserRegistered) error
	GetUserRCount(ctx context.Context) int64
	GetUserRCountBySe(ctx context.Context, start, end uint64) int64
	GetMintNftCountBySe(ctx context.Context, start, end uint64) int64
	GetMintNftUsdtPaidSumBySe(ctx context.Context, start, end uint64) string
	GetMintNftCount(paidType uint64) int64
	GetMintNftNotOpenCount(paidType uint64) int64
	GetMintNftUsdtPaidSum(paidType uint64) string
	GetMintNftNotOpenUsdtPaidSum(paidType uint64) string
	GetNftBuyCountBySe(ctx context.Context, start, end uint64) int64
	GetNftBuySumBySe(ctx context.Context, start, end uint64) string
	GetNftBuySum() string
	GetNftBuyCount() int64
	GetNftOpenCountBySe(ctx context.Context, start, end uint64) int64
	GetNftOpenSum() string
	GetBindReferralLast(ctx context.Context) (*BindReferral, error)
	GetBindReferrals(ctx context.Context) ([]*BindReferral, error)
	InsertBindReferral(ctx context.Context, iData *BindReferral) error
	GetUserV1BoundLast(ctx context.Context) (*UserV1Bound, error)
	GetUserV1BoundByAddress(ctx context.Context, address string) (*UserV1Bound, error)
	GetUserV1Bounds(ctx context.Context) ([]*UserV1Bound, error)
	GetUserV1BoundsByIDs(ctx context.Context, ids []uint64) ([]*UserV1Bound, error)
	GetUserV1Overview(ctx context.Context, yesterdayStart, todayStart, tomorrowStart uint64) (*UserV1Overview, error)
	GetUserV1BoundPage(ctx context.Context, page, pageSize uint64, minAmount, minChildrenAmount, orderBy, order, address string, userID uint64) ([]*UserV1Bound, uint64, error)
	UpdateUserV1Name(ctx context.Context, address, name string) error
	InsertUserV1Bound(ctx context.Context, iData *UserV1Bound) error
	DeleteUserV1BoundAll(ctx context.Context) error
	GetUserV1BoundSyncProgress(ctx context.Context) (*UserV1BoundSyncProgress, error)
	SaveUserV1BoundSyncProgress(ctx context.Context, lastProcessedBlock uint64) error
	GetUserV1PerformanceSyncProgress(ctx context.Context, streamName string) (*UserV1PerformanceSyncProgress, error)
	SaveUserV1PerformanceSyncProgress(ctx context.Context, streamName string, lastProcessedBlock uint64) error
	DeleteUserV1PerformanceSyncProgress(ctx context.Context, streamName string) error
	InsertUserV1StakeChanged(ctx context.Context, event *UserV1StakeChanged) (bool, error)
	InsertUserV1ExtraChanged(ctx context.Context, event *UserV1ExtraChanged) (bool, error)
	UpdateUserV1ExtraChangedApplyStatus(ctx context.Context, eventID uint64, applyStatus uint8) error
	InsertStakingV1TeamBooked(ctx context.Context, event *StakingV1Reward) (bool, error)
	InsertStakingV1TeamClaimed(ctx context.Context, event *StakingV1Reward) (bool, error)
	InsertStakingV1TeamExpired(ctx context.Context, event *StakingV1Reward) (bool, error)
	InsertStakingV1LineClaimed(ctx context.Context, event *StakingV1Reward) (bool, error)
	GetUserV1StakeChangedBlocksWithoutTime(ctx context.Context, fromBlock, limit uint64) ([]uint64, error)
	CountUserV1StakeChangedBlocksWithoutTime(ctx context.Context, fromBlock uint64) (uint64, error)
	UpdateUserV1StakeChangedBlockTimes(ctx context.Context, blockTimes map[uint64]uint64) error
	GetUserV1StakeAddEvents(ctx context.Context) ([]*UserV1StakeChanged, error)
	UpdateUserV1StakeInvestmentNumber(ctx context.Context, eventID, investmentNumber uint64) error
	RepairUserV1InvestmentCount(ctx context.Context) (uint64, uint64, error)
	UpdateUserV1StakeAmount(ctx context.Context, userID uint64, amount string, isAdd bool) error
	UpdateUserV1ChildrenAmount(ctx context.Context, userIDs []uint64, amount string, isAdd bool) error
	UpdateUserV1ExtraAmount(ctx context.Context, userID uint64, amount string) error
	UpdateUserV1TeamBooked(ctx context.Context, userID uint64, event *StakingV1Reward) error
	UpdateUserV1TeamClaimed(ctx context.Context, userID uint64, event *StakingV1Reward) error
	UpdateUserV1TeamExpired(ctx context.Context, userID uint64, amount string) error
	UpdateUserV1LineClaimed(ctx context.Context, userID uint64, event *StakingV1Reward) error
	UpdateUserV1LevelReward(ctx context.Context, userID uint64, amount string) error
	InsertStakingV1OrderEvent(ctx context.Context, event *StakingV1OrderEvent) (bool, error)
	ApplyStakingV1OrderEvent(ctx context.Context, event *StakingV1OrderEvent) error
	ApplyStakingV1OrderSnapshot(ctx context.Context, snapshot *StakingV1OrderSnapshot) error
	MarkStakingV1OrderUsersForSnapshot(ctx context.Context, users []*StakingV1OrderUser) error
	GetStakingV1OrderByOrderID(ctx context.Context, orderID string) (*StakingV1Order, error)
	GetActiveStakingV1OrdersByAddress(ctx context.Context, address string) ([]*StakingV1Order, error)
	GetStakingV1OrderPage(ctx context.Context, query *StakingV1OrderQuery) ([]*StakingV1Order, uint64, error)
	GetStakingV1OrderUsersNeedingSnapshot(ctx context.Context, limit uint64) ([]*StakingV1OrderUser, error)
	CountStakingV1OrderUsersNeedingSnapshot(ctx context.Context) (uint64, error)
	GetStakingV1PlanDaysCounts(ctx context.Context) (map[string]uint32, error)
	RepairStakingV1OrderLinePaid(ctx context.Context) (uint64, error)
	GetStakingV1OrderIntegrity(ctx context.Context) (*StakingV1OrderIntegrity, error)
	IncrementExitedStakingV1OrderLinePaid(ctx context.Context, orderID, grossU string, eventBlock uint64) error
	ResetUserV1Performance(ctx context.Context) error
	DeleteUserV1PerformanceEvents(ctx context.Context) error
	GetStakingStakedLast(ctx context.Context) (*StakingStaked, error)
	InsertStakingStaked(ctx context.Context, iData *StakingStaked) error
	GetStakingQueueAddedLast(ctx context.Context) (*StakingQueueAdded, error)
	InsertStakingQueueAdded(ctx context.Context, iData *StakingQueueAdded) error
}

// AppUsecase is an app usecase.
type AppUsecase struct {
	userRepo UserRepo
	tx       Transaction
	log      *log.Helper
}

// NewAppUsecase new a app usecase.
func NewAppUsecase(userRepo UserRepo, tx Transaction, logger log.Logger) *AppUsecase {
	return &AppUsecase{userRepo: userRepo, tx: tx, log: log.NewHelper(logger)}
}

func (ac *AppUsecase) GetSwapTradeLast(ctx context.Context) (*SwapTrade, error) {
	var (
		rLast *SwapTrade
		err   error
	)
	rLast, err = ac.userRepo.GetSwapTradeLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertSwapTrade(ctx context.Context, trade *SwapTrade) error {
	var (
		err error
	)

	err = ac.userRepo.InsertSwapTrade(ctx, trade)

	return err
}

func (ac *AppUsecase) GetBuyLast(ctx context.Context) (*PrimaryBuy, error) {
	var (
		rLast *PrimaryBuy
		err   error
	)
	rLast, err = ac.userRepo.GetPrimaryBuyLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertBuyTrade(ctx context.Context, trade *PrimaryBuy) error {
	var (
		err error
	)

	err = ac.userRepo.InsertPrimaryBuy(ctx, trade)

	return err
}

func (ac *AppUsecase) GetSellLast(ctx context.Context) (*PrimarySell, error) {
	var (
		rLast *PrimarySell
		err   error
	)
	rLast, err = ac.userRepo.GetPrimarySellLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertSellTrade(ctx context.Context, trade *PrimarySell) error {
	var (
		err error
	)

	err = ac.userRepo.InsertPrimarySell(ctx, trade)

	return err
}

func (ac *AppUsecase) GetRewardLast(ctx context.Context) (*RewardNotified, error) {
	var (
		rLast *RewardNotified
		err   error
	)
	rLast, err = ac.userRepo.GetRewardNotifiedLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertReward(ctx context.Context, trade *RewardNotified) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertRewardNotified(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "分红写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetNftMarketPurchaseLast(ctx context.Context) (*NftMarketPurchase, error) {
	var (
		rLast *NftMarketPurchase
		err   error
	)
	rLast, err = ac.userRepo.GetNftMarketPurchaseLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertNftMarketPurchase(ctx context.Context, trade *NftMarketPurchase) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertNftMarketPurchase(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "买盲盒写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetNftMarketListedLast(ctx context.Context) (*NftMarketListed, error) {
	var (
		rLast *NftMarketListed
		err   error
	)
	rLast, err = ac.userRepo.GetNftMarketListedLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertNftMarketListed(ctx context.Context, trade *NftMarketListed) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertNftMarketListed(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "买盲盒写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetNftMarketUnlistedLast(ctx context.Context) (*NftMarketUnlisted, error) {
	var (
		rLast *NftMarketUnlisted
		err   error
	)
	rLast, err = ac.userRepo.GetNftMarketUnlistedLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertNftMarketUnlisted(ctx context.Context, trade *NftMarketUnlisted) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertNftMarketUnlisted(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "买盲盒写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetNftTransferLast(ctx context.Context) (*NftTransfer, error) {
	var (
		rLast *NftTransfer
		err   error
	)
	rLast, err = ac.userRepo.GetNftTransferLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertNftTransfer(ctx context.Context, trade *NftTransfer) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertNftTransfer(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "买盲盒写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetNftOpenedLast(ctx context.Context) (*NftOpened, error) {
	var (
		rLast *NftOpened
		err   error
	)
	rLast, err = ac.userRepo.GetNftOpenedLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertNftOpened(ctx context.Context, trade *NftOpened) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertNftOpened(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "买盲盒写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetNftMintedLast(ctx context.Context) (*NftMinted, error) {
	var (
		rLast *NftMinted
		err   error
	)
	rLast, err = ac.userRepo.GetNftMintedLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) UpdateNftMinted(ctx context.Context, tokenId uint64, mintAddr string) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.UpdateNftMinted(ctx, tokenId, mintAddr)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "买盲盒写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) InsertNftMinted(ctx context.Context, trade *NftMinted) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertNftMinted(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "买盲盒写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetUserRegisteredLast(ctx context.Context) (*UserRegistered, error) {
	var (
		rLast *UserRegistered
		err   error
	)
	rLast, err = ac.userRepo.GetUserRegisteredLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertUserRegistered(ctx context.Context, trade *UserRegistered) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertUserRegistered(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "注册写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetExchangeList(ctx context.Context, req *pb.GetExchangeListRequest) (*pb.GetExchangeListReply, error) {
	res := make([]*pb.GetExchangeListReply_List, 0)

	var (
		swapTrades []*SwapTrade
		err        error
	)

	swapTrades, err = ac.userRepo.GetSwapTrade(ctx, req.Start, req.End)
	if err != nil {
		return nil, err
	}

	for _, v := range swapTrades {

		var tmpPrice float64
		if 1 == v.Side && 0 < v.Amount0OutGross {
			tmpPrice = v.Amount1In / v.Amount0OutGross
		} else if 2 == v.Side && 0 < v.Amount0In {
			tmpPrice = v.Amount1OutGross / v.Amount0In
		} else {
			continue
		}

		res = append(res, &pb.GetExchangeListReply_List{
			BlockTime: v.BlockTime,
			Price:     tmpPrice,
		})
	}

	return &pb.GetExchangeListReply{List: res}, nil
}

func (ac *AppUsecase) GetBuyList(ctx context.Context, req *pb.GetBuyListRequest) (*pb.GetBuyListReply, error) {
	res := make([]*pb.GetBuyListReply_List, 0)

	var (
		buyList []*PrimaryBuy
		err     error
	)

	buyList, err = ac.userRepo.GetPrimaryBuy(ctx, req.Start, req.End)
	if err != nil {
		return nil, err
	}

	for _, v := range buyList {
		res = append(res, &pb.GetBuyListReply_List{
			BlockTime:    v.BlockTime,
			Price:        v.PriceAfter,
			UsdtUse:      v.UsdtUsed,
			AusdGrossOut: v.AusdGrossOut,
		})
	}

	return &pb.GetBuyListReply{List: res}, nil
}

type Pagination struct {
	PageNum  int
	PageSize int
}

func (ac *AppUsecase) GetRewardList(ctx context.Context, req *pb.GetRewardListRequest) (*pb.GetRewardListReply, error) {
	res := make([]*pb.GetRewardListReply_List, 0)

	var (
		userRewards []*RewardDetail
		count       int64
		err         error
	)

	if 0 > req.Reason || 5 < req.Reason || 0 >= len(req.Address) {
		return &pb.GetRewardListReply{
			Count: 0,
			List:  res,
		}, nil
	}

	userRewards, err, count = ac.userRepo.GetUserRewardByUserIdPage(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	}, req.Address, req.Reason)
	if nil != err {
		return &pb.GetRewardListReply{
			Count: uint64(count),
			List:  res,
		}, err
	}

	tmpIds := make([]uint64, 0)
	for _, vUserReward := range userRewards {
		tmpIds = append(tmpIds, vUserReward.NotifiedId)
	}

	var (
		rn map[uint64]*RewardNotified
	)

	rn, err = ac.userRepo.GetRewardNotifiedByIds(ctx, tmpIds)
	if nil != err {
		return &pb.GetRewardListReply{
			Count: uint64(count),
			List:  res,
		}, err
	}

	for _, vUserReward := range userRewards {
		if _, ok := rn[vUserReward.NotifiedId]; !ok {
			continue
		}

		res = append(res, &pb.GetRewardListReply_List{
			User:      rn[vUserReward.NotifiedId].User,
			Reason:    vUserReward.Reason,
			Amount:    vUserReward.Amount,
			BlockTime: vUserReward.BlockTime,
		})
	}

	return &pb.GetRewardListReply{
		Count: uint64(count),
		List:  res,
	}, nil
}

func (ac *AppUsecase) GetBuyBoxList(ctx context.Context, req *pb.GetBuyBoxListRequest) (*pb.GetBuyBoxListReply, error) {
	res := make([]*pb.GetBuyBoxListReply_List, 0)

	var (
		records []*NftMarketPurchase
		count   int64
		err     error
	)

	records, err, count = ac.userRepo.GetNftMarketPurchaseByAddressPage(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	}, req.Address, nil, 1)
	if nil != err {
		return &pb.GetBuyBoxListReply{
			Count: uint64(count),
			List:  res,
		}, err
	}

	for _, v := range records {
		res = append(res, &pb.GetBuyBoxListReply_List{
			User:      v.Buyer,
			TokenId:   v.TokenID,
			PriceUsdt: v.PriceUSDT,
			FeeUsdt:   v.FeeUSDT,
			FeeB:      v.FeeB,
			BlockTime: v.BlockTime,
		})
	}

	return &pb.GetBuyBoxListReply{
		Count: uint64(count),
		List:  res,
	}, nil
}

func (ac *AppUsecase) UpdateBox(ctx context.Context, req *pb.UpdateBoxRequest) {

	var (
		res       []*NftTransfer
		resTwo    []*NftMarketListed
		resThree  []*NftMarketUnlisted
		resFour   []*NftMarketPurchase
		resFive   []*NftOpened
		resMinted map[uint64]*NftMinted
		err       error
	)

	tokenIdsMap := make(map[uint64]uint64, 0)

	// 交易
	res, err = ac.userRepo.GetNftTransferLastNoCheck(ctx)
	if nil != err {
		return
	}

	for _, v := range res {
		tokenIdsMap[v.TokenID] = v.TokenID
	}

	// 上架
	resTwo, err = ac.userRepo.GetNftListLastNoCheck(ctx)
	if nil != err {
		return
	}

	for _, v := range resTwo {
		tokenIdsMap[v.TokenID] = v.TokenID
	}

	// 下
	resThree, err = ac.userRepo.GetNftUnListLastNoCheck(ctx)
	if nil != err {
		return
	}

	for _, v := range resThree {
		tokenIdsMap[v.TokenID] = v.TokenID
	}

	// 买
	resFour, err = ac.userRepo.GetNftBuyLastNoCheck(ctx)
	if nil != err {
		return
	}

	for _, v := range resFour {
		tokenIdsMap[v.TokenID] = v.TokenID
	}

	// 开
	resFive, err = ac.userRepo.GetNftOpenLastNoCheck(ctx)
	if nil != err {
		return
	}

	for _, v := range resFive {
		tokenIdsMap[v.TokenID] = v.TokenID
	}

	// 整合
	tokenIds := make([]uint64, 0)
	for _, v := range tokenIdsMap {
		tokenIds = append(tokenIds, v)
	}
	if 0 >= len(tokenIds) {
		return
	}

	resMinted, err = ac.userRepo.GetNftMintedByTokenIds(ctx, tokenIds)
	if nil != err || 0 >= len(resMinted) {
		return
	}

	for _, v := range res {
		if _, ok := resMinted[v.TokenID]; !ok {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].BlockTime {
			continue
		}

		//if v.BlockTime <= resMinted[v.TokenID].CheckTime {
		//	fmt.Println("不是最新状态", v, resMinted[v.TokenID])
		//	continue
		//}

		// 上级不改变持有人
		tmpCheck := uint64(1)
		if "0xE447b0391d3F03befeC0dC09E25c049132618fd9" == v.ToAddr {
			tmpCheck = 0
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.UpdateNftMintedToAddress(ctx, resMinted[v.TokenID].ID, v.ID, tmpCheck, v.ToAddr)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "写入mysql错误")
			return
		}

	}

	for _, v := range resTwo {
		if _, ok := resMinted[v.TokenID]; !ok {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].BlockTime {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].CheckTime {
			ac.userRepo.UpdateListedCheckStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			fmt.Println("不是最新状态2", v, resMinted[v.TokenID])
			continue
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.UpdateNftMintedListStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "写入mysql错误")
			ac.userRepo.UpdateListedCheckStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			return
		}
	}

	for _, v := range resThree {
		if _, ok := resMinted[v.TokenID]; !ok {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].BlockTime {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].CheckTime {
			ac.userRepo.UpdateUnlistedCheckStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			fmt.Println("不是最新状态3", v, resMinted[v.TokenID])
			continue
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.UpdateNftMintedUnListStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			ac.userRepo.UpdateUnlistedCheckStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			fmt.Println(err, "写入mysql错误")
			return
		}
	}

	for _, v := range resFour {
		if _, ok := resMinted[v.TokenID]; !ok {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].BlockTime {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].CheckTime {
			ac.userRepo.UpdateBuyCheckStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			fmt.Println("不是最新状态4", v, resMinted[v.TokenID])
			continue
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.UpdateNftMintedBuyStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			ac.userRepo.UpdateBuyCheckStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			fmt.Println(err, "写入mysql错误")
			return
		}
	}

	for _, v := range resFive {
		if _, ok := resMinted[v.TokenID]; !ok {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].BlockTime {
			continue
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.UpdateNftMintedOpenStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "写入mysql错误")
			return
		}
	}
}

func (ac *AppUsecase) GetAddressBox(ctx context.Context, address []string, req *pb.GetAddressBoxRequest) ([]*NftMinted, int64, error) {
	var (
		minted []*NftMinted
		count  int64
		err    error
	)

	minted, err, count = ac.userRepo.GetNftMintedByAddressPage(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, req.Start, req.End, address, req.Num, req.NumThree, req.NumTwo, req.OpenStatus, req.NumFour)

	return minted, count, err
}

func (ac *AppUsecase) GetNftMintedPage(ctx context.Context, req *pb.GetMarketListRequest) ([]*NftMinted, int64, error) {
	var (
		minted []*NftMinted
		count  int64
		err    error
	)

	minted, err, count = ac.userRepo.GetNftMintedPage(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, req.Num, req.NumThree, req.NumTwo)

	return minted, count, err
}

func (ac *AppUsecase) GetSellBoxList(ctx context.Context, address []string, req *pb.GetSellBoxListRequest) ([]*NftMarketPurchase, int64, error) {
	var (
		records []*NftMarketPurchase
		count   int64
		err     error
	)

	records, err, count = ac.userRepo.GetNftMarketPurchaseByAddressPage(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, req.Address, address, 3)

	return records, count, err
}

func (ac *AppUsecase) GetAllInfo(ctx context.Context, req *pb.GetAllInfoRequest) (*pb.GetAllInfoReply, error) {

	return &pb.GetAllInfoReply{
		UserCount:         uint64(ac.userRepo.GetUserRCount(ctx)),
		TodayUserCount:    uint64(ac.userRepo.GetUserRCountBySe(ctx, req.Start, req.End)),
		TodayMintedCount:  uint64(ac.userRepo.GetMintNftCountBySe(ctx, req.Start, req.End)),
		TodayMintedSum:    ac.userRepo.GetMintNftUsdtPaidSumBySe(ctx, req.Start, req.End),
		MintedCount:       uint64(ac.userRepo.GetMintNftCount(0)),
		MintedSum:         ac.userRepo.GetMintNftUsdtPaidSum(0),
		MintedCountOne:    uint64(ac.userRepo.GetMintNftCount(1)),
		MintedSumOne:      ac.userRepo.GetMintNftUsdtPaidSum(1),
		MintedCountTwo:    uint64(ac.userRepo.GetMintNftCount(2)),
		MintedSumTwo:      ac.userRepo.GetMintNftUsdtPaidSum(2),
		MintedCountThree:  uint64(ac.userRepo.GetMintNftCount(3)),
		MintedSumThree:    ac.userRepo.GetMintNftUsdtPaidSum(3),
		BuyCount:          uint64(ac.userRepo.GetNftBuyCount()),
		BuySum:            ac.userRepo.GetNftBuySum(),
		TodayBuyCount:     uint64(ac.userRepo.GetNftBuyCountBySe(ctx, req.Start, req.End)),
		TodayBuySum:       ac.userRepo.GetNftBuySumBySe(ctx, req.Start, req.End),
		OpenReward:        ac.userRepo.GetNftOpenSum(),
		MintedCountNoOpen: uint64(ac.userRepo.GetMintNftNotOpenCount(0)),
		MintedSumNoOpen:   ac.userRepo.GetMintNftNotOpenUsdtPaidSum(0),
	}, nil
}

func (ac *AppUsecase) GetBindReferralLast(ctx context.Context) (*BindReferral, error) {
	var (
		rLast *BindReferral
		err   error
	)
	rLast, err = ac.userRepo.GetBindReferralLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) GetBindReferrals(ctx context.Context) ([]*BindReferral, error) {
	var (
		rLast []*BindReferral
		err   error
	)
	rLast, err = ac.userRepo.GetBindReferrals(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertBindReferral(ctx context.Context, trade *BindReferral) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertBindReferral(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "bind写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetUserV1BoundLast(ctx context.Context) (*UserV1Bound, error) {
	var (
		rLast *UserV1Bound
		err   error
	)
	rLast, err = ac.userRepo.GetUserV1BoundLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertUserV1Bound(ctx context.Context, trade *UserV1Bound) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertUserV1Bound(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "bind写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetUserV1BoundSyncProgress(ctx context.Context) (*UserV1BoundSyncProgress, error) {
	progress, err := ac.userRepo.GetUserV1BoundSyncProgress(ctx)
	if nil != err || nil == progress {
		return nil, err
	}

	return progress, nil
}

func (ac *AppUsecase) SaveUserV1BoundRange(ctx context.Context, events []*UserV1Bound, lastProcessedBlock uint64) error {
	var err error

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		for _, event := range events {
			existing, errT := ac.userRepo.GetUserV1BoundByAddress(ctx, event.UserAddr)
			if nil != errT {
				return errT
			}
			if nil != existing {
				if existing.BlockNumber != event.BlockNumber || !strings.EqualFold(existing.ParentAddr, event.ParentAddr) {
					return fmt.Errorf("user %s event mismatch: db_block=%d chain_block=%d db_parent=%s chain_parent=%s", event.UserAddr, existing.BlockNumber, event.BlockNumber, existing.ParentAddr, event.ParentAddr)
				}
				continue
			}

			var parent *UserV1Bound
			if strings.EqualFold(event.ParentAddr, "0x0000000000000000000000000000000000000000") {
				event.RecommendCode = ""
			} else {
				parent, errT = ac.userRepo.GetUserV1BoundByAddress(ctx, event.ParentAddr)
				if nil != errT {
					return errT
				}
				if nil == parent {
					return fmt.Errorf("user %s parent %s does not exist", event.UserAddr, event.ParentAddr)
				}
				event.RecommendCode = parent.RecommendCode + "D" + strconv.FormatUint(parent.ID, 10)
			}

			if err = ac.userRepo.InsertUserV1Bound(ctx, event); nil != err {
				return err
			}
			if nil != parent && parent.ID >= event.ID {
				return fmt.Errorf("user %s id %d must be greater than parent id %d", event.UserAddr, event.ID, parent.ID)
			}
		}

		if err = ac.userRepo.SaveUserV1BoundSyncProgress(ctx, lastProcessedBlock); nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "user bound分段写入mysql错误")
		return err
	}

	return nil
}

func (ac *AppUsecase) RebuildUserV1Bound(ctx context.Context, events []*UserV1Bound, lastProcessedBlock uint64) error {
	var err error

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		existingRows, errT := ac.userRepo.GetUserV1Bounds(ctx)
		if nil != errT {
			return errT
		}
		existingByAddress := make(map[string]*UserV1Bound, len(existingRows))
		for _, row := range existingRows {
			existingByAddress[strings.ToLower(row.UserAddr)] = row
		}
		for _, event := range events {
			if existing := existingByAddress[strings.ToLower(event.UserAddr)]; nil != existing {
				copyUserV1Performance(event, existing)
			}
		}

		if err = ac.userRepo.DeleteUserV1BoundAll(ctx); nil != err {
			return err
		}

		for _, event := range events {
			if err = ac.userRepo.InsertUserV1Bound(ctx, event); nil != err {
				return err
			}
		}

		if err = ac.userRepo.SaveUserV1BoundSyncProgress(ctx, lastProcessedBlock); nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "user bound历史重建写入mysql错误")
		return err
	}

	return nil
}

func copyUserV1Performance(target *UserV1Bound, source *UserV1Bound) {
	target.Name = source.Name
	target.Amount = source.Amount
	target.AmountHistory = source.AmountHistory
	target.InvestmentCount = source.InvestmentCount
	target.ChildrenAmount = source.ChildrenAmount
	target.ChildrenAmountHistory = source.ChildrenAmountHistory
	target.ChildrenAmountExtra = source.ChildrenAmountExtra
	target.RewardRecommendAmount = source.RewardRecommendAmount
	target.RewardRecommendPay = source.RewardRecommendPay
	target.RewardRecommendStoreAmount = source.RewardRecommendStoreAmount
	target.RewardRecommendFee = source.RewardRecommendFee
	target.RewardRecommendTeamUAmount = source.RewardRecommendTeamUAmount
	target.RewardRecommendClaimedTeamUNet = source.RewardRecommendClaimedTeamUNet
	target.RewardRecommendClaimedTeamUAmount = source.RewardRecommendClaimedTeamUAmount
	target.RewardRecommendClaimedTeamUFee = source.RewardRecommendClaimedTeamUFee
	target.RewardRecommendExpired = source.RewardRecommendExpired
	target.LineU = source.LineU
	target.LineCoinU = source.LineCoinU
	target.LineCoin = source.LineCoin
	target.LineFee = source.LineFee
	target.LevelReward = source.LevelReward
}

func (ac *AppUsecase) GetUserV1BoundByAddress(ctx context.Context, address string) (*UserV1Bound, error) {
	return ac.userRepo.GetUserV1BoundByAddress(ctx, strings.ToLower(address))
}

func (ac *AppUsecase) GetUserV1Bounds(ctx context.Context) ([]*UserV1Bound, error) {
	return ac.userRepo.GetUserV1Bounds(ctx)
}

func (ac *AppUsecase) GetUserV1Overview(ctx context.Context, yesterdayStart, todayStart, tomorrowStart uint64) (*UserV1Overview, error) {
	overview, err := ac.userRepo.GetUserV1Overview(ctx, yesterdayStart, todayStart, tomorrowStart)
	if nil != err {
		return nil, err
	}
	if 0 < overview.MissingInvestmentBlockTimeEventCount {
		return nil, fmt.Errorf("有 %d 条投资事件缺少链上时间，请先执行 /api/recover_user_investment_data", overview.MissingInvestmentBlockTimeEventCount)
	}
	if 0 < overview.MissingInvestmentNumberEventCount {
		return nil, fmt.Errorf("有 %d 条投资事件缺少投资次序，请继续执行 /api/recover_user_investment_data", overview.MissingInvestmentNumberEventCount)
	}
	return overview, nil
}

func (ac *AppUsecase) GetUserV1BoundPage(ctx context.Context, page, pageSize uint64, minAmount, minChildrenAmount, orderBy, order, address string, userID uint64) ([]*UserV1Bound, uint64, error) {
	return ac.userRepo.GetUserV1BoundPage(ctx, page, pageSize, minAmount, minChildrenAmount, orderBy, order, address, userID)
}

func (ac *AppUsecase) UpdateUserV1Name(ctx context.Context, address, name string) error {
	return ac.userRepo.UpdateUserV1Name(ctx, address, name)
}

func (ac *AppUsecase) GetUserV1StakeChangedBlocksWithoutTime(ctx context.Context, fromBlock, limit uint64) ([]uint64, error) {
	return ac.userRepo.GetUserV1StakeChangedBlocksWithoutTime(ctx, fromBlock, limit)
}

func (ac *AppUsecase) CountUserV1StakeChangedBlocksWithoutTime(ctx context.Context, fromBlock uint64) (uint64, error) {
	return ac.userRepo.CountUserV1StakeChangedBlocksWithoutTime(ctx, fromBlock)
}

func (ac *AppUsecase) SaveUserV1StakeChangedBlockTimes(ctx context.Context, blockTimes map[uint64]uint64) error {
	return ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		return ac.userRepo.UpdateUserV1StakeChangedBlockTimes(ctx, blockTimes)
	})
}

func (ac *AppUsecase) RepairUserV1InvestmentData(ctx context.Context) (uint64, uint64, error) {
	progress, err := ac.userRepo.GetUserV1PerformanceSyncProgress(ctx, UserV1PerformanceStreamStake)
	if nil != err {
		return 0, 0, err
	}
	if nil == progress {
		return 0, 0, fmt.Errorf("StakeChanged 同步进度不存在，请先完成 /api/recover_performance_event")
	}

	var userCount uint64
	var orderCount uint64
	err = ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		events, errT := ac.userRepo.GetUserV1StakeAddEvents(ctx)
		if nil != errT {
			return errT
		}
		lastUser := ""
		investmentNumber := uint64(0)
		for _, event := range events {
			if event.UserAddr != lastUser {
				lastUser = event.UserAddr
				investmentNumber = 0
			}
			investmentNumber++
			if errT = ac.userRepo.UpdateUserV1StakeInvestmentNumber(ctx, event.ID, investmentNumber); nil != errT {
				return errT
			}
		}

		userCount, orderCount, errT = ac.userRepo.RepairUserV1InvestmentCount(ctx)
		return errT
	})
	return userCount, orderCount, err
}

func (ac *AppUsecase) GetUserV1PerformanceSyncProgress(ctx context.Context, streamName string) (*UserV1PerformanceSyncProgress, error) {
	return ac.userRepo.GetUserV1PerformanceSyncProgress(ctx, streamName)
}

func userV1RecommendIDs(recommendCode string) ([]uint64, error) {
	recommendCode = strings.TrimSpace(recommendCode)
	if "" == recommendCode {
		return []uint64{}, nil
	}
	if !strings.HasPrefix(recommendCode, "D") {
		return nil, fmt.Errorf("bad recommend code %q", recommendCode)
	}

	parts := strings.Split(recommendCode, "D")
	ids := make([]uint64, 0, len(parts)-1)
	for _, part := range parts[1:] {
		if "" == part {
			return nil, fmt.Errorf("bad recommend code %q", recommendCode)
		}
		id, err := strconv.ParseUint(part, 10, 64)
		if nil != err || 0 == id {
			return nil, fmt.Errorf("bad recommend user id %q in %q", part, recommendCode)
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (ac *AppUsecase) checkedUserV1Ancestors(ctx context.Context, user *UserV1Bound, includeRootSelf bool) ([]uint64, error) {
	ids, err := userV1RecommendIDs(user.RecommendCode)
	if nil != err {
		return nil, err
	}
	if includeRootSelf && 1 == user.ID && 0 == len(ids) {
		// UserV1 的根用户在 _walk(root) 中会把自己的质押计入 basePerf。
		ids = append(ids, user.ID)
	}
	if 0 == len(ids) {
		return ids, nil
	}
	rows, err := ac.userRepo.GetUserV1BoundsByIDs(ctx, ids)
	if nil != err {
		return nil, err
	}
	if len(rows) != len(ids) {
		return nil, fmt.Errorf("user %s recommend path is incomplete: code=%s ids=%d rows=%d", user.UserAddr, user.RecommendCode, len(ids), len(rows))
	}
	return ids, nil
}

func (ac *AppUsecase) saveUserV1StakeChangedEvents(ctx context.Context, events []*UserV1StakeChanged) error {
	for _, event := range events {
		inserted, err := ac.userRepo.InsertUserV1StakeChanged(ctx, event)
		if nil != err {
			return err
		}
		if !inserted {
			continue
		}

		user, err := ac.userRepo.GetUserV1BoundByAddress(ctx, strings.ToLower(event.UserAddr))
		if nil != err {
			return err
		}
		if nil == user {
			return fmt.Errorf("StakeChanged user %s is not in user_v1_bound_event", event.UserAddr)
		}
		ancestorIDs, err := ac.checkedUserV1Ancestors(ctx, user, true)
		if nil != err {
			return err
		}
		if event.IsAdd {
			event.InvestmentNumber = user.InvestmentCount + 1
			if err = ac.userRepo.UpdateUserV1StakeInvestmentNumber(ctx, event.ID, event.InvestmentNumber); nil != err {
				return err
			}
		}
		if err = ac.userRepo.UpdateUserV1StakeAmount(ctx, user.ID, event.Amount, event.IsAdd); nil != err {
			return err
		}
		if err = ac.userRepo.UpdateUserV1ChildrenAmount(ctx, ancestorIDs, event.Amount, event.IsAdd); nil != err {
			return err
		}
	}
	return nil
}

func (ac *AppUsecase) saveUserV1ExtraChangedEvents(ctx context.Context, events []*UserV1ExtraChanged) error {
	for _, event := range events {
		inserted, err := ac.userRepo.InsertUserV1ExtraChanged(ctx, event)
		if nil != err {
			return err
		}
		if !inserted {
			continue
		}

		user, err := ac.userRepo.GetUserV1BoundByAddress(ctx, strings.ToLower(event.UserAddr))
		if nil != err {
			return err
		}
		if nil == user {
			if err = ac.userRepo.UpdateUserV1ExtraChangedApplyStatus(ctx, event.ID, UserV1ExtraChangedApplyStatusUnregistered); nil != err {
				return err
			}
			fmt.Printf("ExtraChanged event marked unregistered_pending address=%s block=%d event_key=%s\n", event.UserAddr, event.BlockNumber, event.EventKey)
			continue
		}
		if err = ac.userRepo.UpdateUserV1ExtraAmount(ctx, user.ID, event.ExtraAmount); nil != err {
			return err
		}
	}
	return nil
}

func (ac *AppUsecase) saveStakingV1RewardEvents(ctx context.Context, events []*StakingV1Reward) error {
	for _, event := range events {
		var (
			inserted bool
			err      error
			address  string
		)
		switch event.EventType {
		case StakingV1RewardTeamBooked:
			inserted, err = ac.userRepo.InsertStakingV1TeamBooked(ctx, event)
			address = event.ToAddr
		case StakingV1RewardTeamClaimed:
			inserted, err = ac.userRepo.InsertStakingV1TeamClaimed(ctx, event)
			address = event.UserAddr
		case StakingV1RewardTeamExpired:
			inserted, err = ac.userRepo.InsertStakingV1TeamExpired(ctx, event)
			address = event.ToAddr
		case StakingV1RewardLineClaimed:
			inserted, err = ac.userRepo.InsertStakingV1LineClaimed(ctx, event)
			address = event.UserAddr
		default:
			return fmt.Errorf("unknown staking reward event type %q", event.EventType)
		}
		if nil != err {
			return err
		}
		if !inserted {
			continue
		}

		user, err := ac.userRepo.GetUserV1BoundByAddress(ctx, strings.ToLower(address))
		if nil != err {
			return err
		}
		if nil == user {
			return fmt.Errorf("%s user %s is not in user_v1_bound_event", event.EventType, address)
		}

		switch event.EventType {
		case StakingV1RewardTeamBooked:
			err = ac.userRepo.UpdateUserV1TeamBooked(ctx, user.ID, event)
		case StakingV1RewardTeamClaimed:
			err = ac.userRepo.UpdateUserV1TeamClaimed(ctx, user.ID, event)
		case StakingV1RewardTeamExpired:
			err = ac.userRepo.UpdateUserV1TeamExpired(ctx, user.ID, event.Amount)
		case StakingV1RewardLineClaimed:
			err = ac.userRepo.UpdateUserV1LineClaimed(ctx, user.ID, event)
		}
		if nil != err {
			return err
		}
		if StakingV1RewardLineClaimed == event.EventType {
			// Only a newly inserted LineClaimed event reaches here. Increment
			// when the order snapshot is older than the event; a snapshot from
			// the same/newer block already contains this payment. This makes the
			// independently scheduled reward and order endpoints converge in
			// either execution order, including the final payment before exit.
			if err = ac.userRepo.IncrementExitedStakingV1OrderLinePaid(ctx, event.OrderID, event.GrossU, event.BlockNumber); nil != err {
				return err
			}
		}
	}
	return nil
}

func (ac *AppUsecase) checkedStakingV1OrderUser(ctx context.Context, userID uint64, address string) (uint64, string, error) {
	address = strings.ToLower(strings.TrimSpace(address))
	if "" == address {
		return 0, "", fmt.Errorf("staking order user address is empty")
	}
	user, err := ac.userRepo.GetUserV1BoundByAddress(ctx, address)
	if nil != err {
		return 0, "", err
	}
	if nil == user {
		return 0, "", fmt.Errorf("staking order user %s is not in user_v1_bound_event", address)
	}
	if 0 != userID && user.ID != userID {
		return 0, "", fmt.Errorf("staking order user mismatch: address=%s db_user_id=%d event_user_id=%d", address, user.ID, userID)
	}
	return user.ID, address, nil
}

func (ac *AppUsecase) saveStakingV1OrderEvents(ctx context.Context, events []*StakingV1OrderEvent) error {
	planDays, err := ac.userRepo.GetStakingV1PlanDaysCounts(ctx)
	if nil != err {
		return err
	}
	for _, event := range events {
		if nil == event {
			continue
		}
		if "" == strings.TrimSpace(event.EventKey) {
			return fmt.Errorf("%s staking order event has empty event key", event.EventType)
		}
		if StakingV1OrderEventPlanSet == event.EventType {
			inserted, err := ac.userRepo.InsertStakingV1OrderEvent(ctx, event)
			if nil != err {
				return err
			}
			// Events are replayed in block/log order. Keep the in-memory plan
			// state in sync even when this row was already present, so a Created
			// event later in the same range receives the plan's historical days.
			planDays[event.PlanID] = event.DaysCount
			if inserted {
				continue
			}
			continue
		}
		if "" == strings.TrimSpace(event.OrderID) {
			return fmt.Errorf("%s staking order event has empty order id or event key", event.EventType)
		}
		if StakingV1OrderEventCreated == event.EventType {
			daysCount, ok := planDays[event.PlanID]
			if !ok {
				return fmt.Errorf("staking order %s references unknown plan %s", event.OrderID, event.PlanID)
			}
			event.DaysCount = daysCount
		}
		userID, address, err := ac.checkedStakingV1OrderUser(ctx, event.UserID, event.UserAddr)
		if nil != err {
			return err
		}
		event.UserID = userID
		event.UserAddr = address
		inserted, err := ac.userRepo.InsertStakingV1OrderEvent(ctx, event)
		if nil != err {
			return err
		}
		if !inserted {
			continue
		}
		if err = ac.userRepo.ApplyStakingV1OrderEvent(ctx, event); nil != err {
			return err
		}
	}
	return nil
}

func (ac *AppUsecase) applyStakingV1OrderSnapshots(ctx context.Context, snapshots []*StakingV1OrderSnapshot) error {
	for _, snapshot := range snapshots {
		if nil == snapshot {
			continue
		}
		if "" == strings.TrimSpace(snapshot.OrderID) {
			return fmt.Errorf("staking order snapshot has empty order id")
		}
		if snapshot.Status < StakingV1OrderStatusQueued || snapshot.Status > StakingV1OrderStatusExited {
			return fmt.Errorf("staking order %s snapshot has invalid status %d", snapshot.OrderID, snapshot.Status)
		}
		userID, address, err := ac.checkedStakingV1OrderUser(ctx, snapshot.UserID, snapshot.UserAddr)
		if nil != err {
			return err
		}
		snapshot.UserID = userID
		snapshot.UserAddr = address
		if err = ac.userRepo.ApplyStakingV1OrderSnapshot(ctx, snapshot); nil != err {
			return err
		}
	}
	return nil
}

// SaveStakingV1OrderRange atomically stores lifecycle events, marks every
// affected user's active orders dirty, and advances the staking_order
// checkpoint. Dirty orders are refreshed only after the event stream catches
// the recent safe head, so a pruned full node never needs historical state.
func (ac *AppUsecase) SaveStakingV1OrderRange(ctx context.Context, events []*StakingV1OrderEvent, refreshUsers []*StakingV1OrderUser, lastProcessedBlock uint64) error {
	return ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		if err := ac.saveStakingV1OrderEvents(ctx, events); nil != err {
			return err
		}
		checkedUsers := make([]*StakingV1OrderUser, 0, len(refreshUsers))
		for _, user := range refreshUsers {
			if nil == user {
				continue
			}
			userID, address, err := ac.checkedStakingV1OrderUser(ctx, user.UserID, user.UserAddr)
			if nil != err {
				return err
			}
			checkedUsers = append(checkedUsers, &StakingV1OrderUser{UserID: userID, UserAddr: address})
		}
		if err := ac.userRepo.MarkStakingV1OrderUsersForSnapshot(ctx, checkedUsers); nil != err {
			return err
		}
		if err := ac.ValidateStakingV1OrderIntegrity(ctx, ""); nil != err {
			return err
		}
		return ac.userRepo.SaveUserV1PerformanceSyncProgress(ctx, UserV1PerformanceStreamOrder, lastProcessedBlock)
	})
}

// RecoverStakingV1OrderRange only rebuilds event-derived state. Its rows keep
// last_synced_block=0 until the separate snapshot phase reads each active user
// from the chain. The normal staking_order stream is deliberately not opened
// until CompleteStakingV1OrderRecovery succeeds.
func (ac *AppUsecase) RecoverStakingV1OrderRange(ctx context.Context, events []*StakingV1OrderEvent, refreshUsers []*StakingV1OrderUser, lastProcessedBlock uint64) error {
	return ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		if err := ac.saveStakingV1OrderEvents(ctx, events); nil != err {
			return err
		}
		checkedUsers := make([]*StakingV1OrderUser, 0, len(refreshUsers))
		for _, user := range refreshUsers {
			if nil == user {
				continue
			}
			userID, address, err := ac.checkedStakingV1OrderUser(ctx, user.UserID, user.UserAddr)
			if nil != err {
				return err
			}
			checkedUsers = append(checkedUsers, &StakingV1OrderUser{UserID: userID, UserAddr: address})
		}
		if err := ac.userRepo.MarkStakingV1OrderUsersForSnapshot(ctx, checkedUsers); nil != err {
			return err
		}
		if err := ac.ValidateStakingV1OrderIntegrity(ctx, ""); nil != err {
			return err
		}
		return ac.userRepo.SaveUserV1PerformanceSyncProgress(ctx, UserV1PerformanceStreamOrderRecovery, lastProcessedBlock)
	})
}

func (ac *AppUsecase) CompleteStakingV1OrderRecovery(ctx context.Context, lastProcessedBlock uint64) error {
	return ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		for _, streamName := range []string{UserV1PerformanceStreamOrderRecovery, UserV1PerformanceStreamOrder} {
			if err := ac.userRepo.SaveUserV1PerformanceSyncProgress(ctx, streamName, lastProcessedBlock); nil != err {
				return err
			}
		}
		return nil
	})
}

func (ac *AppUsecase) SaveStakingV1OrderRecoveryTarget(ctx context.Context, targetBlock uint64) error {
	return ac.userRepo.SaveUserV1PerformanceSyncProgress(ctx, UserV1PerformanceStreamOrderTarget, targetBlock)
}

func (ac *AppUsecase) ApplyStakingV1OrderSnapshots(ctx context.Context, snapshots []*StakingV1OrderSnapshot) error {
	return ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		return ac.applyStakingV1OrderSnapshots(ctx, snapshots)
	})
}

func (ac *AppUsecase) GetStakingV1OrderByOrderID(ctx context.Context, orderID string) (*StakingV1Order, error) {
	return ac.userRepo.GetStakingV1OrderByOrderID(ctx, strings.TrimSpace(orderID))
}

func (ac *AppUsecase) GetActiveStakingV1OrdersByAddress(ctx context.Context, address string) ([]*StakingV1Order, error) {
	return ac.userRepo.GetActiveStakingV1OrdersByAddress(ctx, strings.ToLower(strings.TrimSpace(address)))
}

func (ac *AppUsecase) GetStakingV1OrderPage(ctx context.Context, query *StakingV1OrderQuery) ([]*StakingV1Order, uint64, error) {
	if nil == query {
		return nil, 0, fmt.Errorf("staking order query is nil")
	}
	query.Address = strings.ToLower(strings.TrimSpace(query.Address))
	query.OrderBy = strings.ToLower(strings.TrimSpace(query.OrderBy))
	query.Order = strings.ToLower(strings.TrimSpace(query.Order))
	return ac.userRepo.GetStakingV1OrderPage(ctx, query)
}

func (ac *AppUsecase) GetStakingV1OrderUsersNeedingSnapshot(ctx context.Context, limit uint64) ([]*StakingV1OrderUser, error) {
	return ac.userRepo.GetStakingV1OrderUsersNeedingSnapshot(ctx, limit)
}

func (ac *AppUsecase) CountStakingV1OrderUsersNeedingSnapshot(ctx context.Context) (uint64, error) {
	return ac.userRepo.CountStakingV1OrderUsersNeedingSnapshot(ctx)
}

func (ac *AppUsecase) GetStakingV1PlanDaysCounts(ctx context.Context) (map[string]uint32, error) {
	return ac.userRepo.GetStakingV1PlanDaysCounts(ctx)
}

func (ac *AppUsecase) RepairStakingV1OrderLinePaid(ctx context.Context) (uint64, error) {
	var repaired uint64
	err := ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		var err error
		repaired, err = ac.userRepo.RepairStakingV1OrderLinePaid(ctx)
		return err
	})
	return repaired, err
}

func (ac *AppUsecase) ValidateStakingV1OrderIntegrity(ctx context.Context, expectedNextOrderID string) error {
	state, err := ac.userRepo.GetStakingV1OrderIntegrity(ctx)
	if nil != err {
		return err
	}
	checks := []struct {
		name  string
		count uint64
	}{
		{"master_without_created", state.MasterWithoutCreated},
		{"created_without_master", state.CreatedWithoutMaster},
		{"exited_without_exit", state.ExitedWithoutExit},
		{"exit_not_marked_exited", state.ExitNotMarkedExited},
		{"running_without_entered", state.RunningWithoutEntered},
		{"queued_with_entered", state.QueuedWithEntered},
		{"queue_done_without_queued", state.QueueDoneWithoutQueued},
		{"queue_done_without_entered", state.QueueDoneWithoutEntered},
		{"lifecycle_identity_mismatch", state.LifecycleIdentityMismatch},
		{"master_created_mismatch", state.MasterCreatedMismatch},
		{"duplicate_created_order_id", state.DuplicateCreatedOrderID},
		{"duplicate_exit_order_id", state.DuplicateExitOrderID},
	}
	for _, check := range checks {
		if 0 != check.count {
			return fmt.Errorf("staking order integrity %s=%d", check.name, check.count)
		}
	}
	if "" == strings.TrimSpace(expectedNextOrderID) {
		return nil
	}
	nextOrderID, ok := new(big.Int).SetString(strings.TrimSpace(expectedNextOrderID), 10)
	if !ok || nextOrderID.Sign() <= 0 {
		return fmt.Errorf("invalid staking nextOrderId %q", expectedNextOrderID)
	}
	expectedCount := new(big.Int).Sub(new(big.Int).Set(nextOrderID), big.NewInt(1))
	actualCount := new(big.Int).SetUint64(state.CreatedCount)
	if actualCount.Cmp(expectedCount) != 0 {
		return fmt.Errorf("staking order created count mismatch: expected=%s actual=%d", expectedCount.String(), state.CreatedCount)
	}
	if 0 == expectedCount.Sign() {
		return nil
	}
	minOrderID, minOK := new(big.Int).SetString(state.MinCreatedOrderID, 10)
	maxOrderID, maxOK := new(big.Int).SetString(state.MaxCreatedOrderID, 10)
	if !minOK || !maxOK || minOrderID.Cmp(big.NewInt(1)) != 0 || maxOrderID.Cmp(expectedCount) != 0 {
		return fmt.Errorf("staking order id range mismatch: expected=1..%s actual=%s..%s", expectedCount.String(), state.MinCreatedOrderID, state.MaxCreatedOrderID)
	}
	return nil
}

// SyncStakingV1LineClaimedOrder refreshes an active order absolutely. For an
// exited order the corresponding newly-inserted LineClaimed event has already
// updated line_paid inside saveStakingV1RewardEvents, so no snapshot may reopen
// or otherwise mutate that final order.
func (ac *AppUsecase) SyncStakingV1LineClaimedOrder(ctx context.Context, event *StakingV1Reward, snapshot *StakingV1OrderSnapshot) error {
	if nil == event {
		return fmt.Errorf("LineClaimed event is nil")
	}
	order, err := ac.userRepo.GetStakingV1OrderByOrderID(ctx, event.OrderID)
	if nil != err {
		return err
	}
	if nil != order && StakingV1OrderStatusExited == order.Status {
		return nil
	}
	if nil == snapshot {
		return fmt.Errorf("active LineClaimed order %s has no chain snapshot", event.OrderID)
	}
	return ac.ApplyStakingV1OrderSnapshots(ctx, []*StakingV1OrderSnapshot{snapshot})
}

func (ac *AppUsecase) SaveUserV1StakeChangedRange(ctx context.Context, events []*UserV1StakeChanged, lastProcessedBlock uint64) error {
	return ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		if err := ac.saveUserV1StakeChangedEvents(ctx, events); nil != err {
			return err
		}
		return ac.userRepo.SaveUserV1PerformanceSyncProgress(ctx, UserV1PerformanceStreamStake, lastProcessedBlock)
	})
}

func (ac *AppUsecase) SaveUserV1ExtraChangedRange(ctx context.Context, events []*UserV1ExtraChanged, lastProcessedBlock uint64) error {
	return ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		if err := ac.saveUserV1ExtraChangedEvents(ctx, events); nil != err {
			return err
		}
		return ac.userRepo.SaveUserV1PerformanceSyncProgress(ctx, UserV1PerformanceStreamExtra, lastProcessedBlock)
	})
}

func (ac *AppUsecase) SaveStakingV1RewardRange(ctx context.Context, events []*StakingV1Reward, levelRewards map[uint64]string, lastProcessedBlock uint64) error {
	return ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		if err := ac.saveStakingV1RewardEvents(ctx, events); nil != err {
			return err
		}
		for userID, amount := range levelRewards {
			if err := ac.userRepo.UpdateUserV1LevelReward(ctx, userID, amount); nil != err {
				return err
			}
		}
		return ac.userRepo.SaveUserV1PerformanceSyncProgress(ctx, UserV1PerformanceStreamReward, lastProcessedBlock)
	})
}

func (ac *AppUsecase) InitializeUserV1PerformanceRecovery(ctx context.Context, startBlock uint64) error {
	return ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		if err := ac.userRepo.ResetUserV1Performance(ctx); nil != err {
			return err
		}
		if err := ac.userRepo.DeleteUserV1PerformanceEvents(ctx); nil != err {
			return err
		}
		for _, streamName := range []string{
			UserV1PerformanceStreamStake, UserV1PerformanceStreamExtra, UserV1PerformanceStreamReward,
		} {
			if err := ac.userRepo.DeleteUserV1PerformanceSyncProgress(ctx, streamName); nil != err {
				return err
			}
		}
		return ac.userRepo.SaveUserV1PerformanceSyncProgress(ctx, UserV1PerformanceStreamRecovery, startBlock)
	})
}

func (ac *AppUsecase) SaveUserV1PerformanceRecoveryRange(ctx context.Context, stakes []*UserV1StakeChanged, extras []*UserV1ExtraChanged, rewards []*StakingV1Reward, lastProcessedBlock uint64) error {
	return ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		if err := ac.saveUserV1StakeChangedEvents(ctx, stakes); nil != err {
			return err
		}
		if err := ac.saveUserV1ExtraChangedEvents(ctx, extras); nil != err {
			return err
		}
		if err := ac.saveStakingV1RewardEvents(ctx, rewards); nil != err {
			return err
		}
		// 历史恢复未完成前只推进 recovery；完成后再一次性开放三个增量接口。
		return ac.userRepo.SaveUserV1PerformanceSyncProgress(ctx, UserV1PerformanceStreamRecovery, lastProcessedBlock)
	})
}

func (ac *AppUsecase) CompleteUserV1PerformanceRecovery(ctx context.Context, levelRewards map[uint64]string, lastProcessedBlock uint64) error {
	return ac.tx.ExecTx(ctx, func(ctx context.Context) error {
		for userID, amount := range levelRewards {
			if err := ac.userRepo.UpdateUserV1LevelReward(ctx, userID, amount); nil != err {
				return err
			}
		}
		for _, streamName := range []string{
			UserV1PerformanceStreamStake, UserV1PerformanceStreamExtra,
			UserV1PerformanceStreamReward, UserV1PerformanceStreamRecovery,
		} {
			if err := ac.userRepo.SaveUserV1PerformanceSyncProgress(ctx, streamName, lastProcessedBlock); nil != err {
				return err
			}
		}
		return nil
	})
}

func (ac *AppUsecase) GetStakingV1LineRewardUsers(ctx context.Context, events []*StakingV1Reward) ([]*UserV1Bound, error) {
	ids := make([]uint64, 0)
	seen := make(map[uint64]struct{})
	for _, event := range events {
		if StakingV1RewardLineClaimed != event.EventType {
			continue
		}
		user, err := ac.userRepo.GetUserV1BoundByAddress(ctx, strings.ToLower(event.UserAddr))
		if nil != err {
			return nil, err
		}
		if nil == user {
			return nil, fmt.Errorf("LineClaimed user %s is not in user_v1_bound_event", event.UserAddr)
		}
		ancestorIDs, err := ac.checkedUserV1Ancestors(ctx, user, false)
		if nil != err {
			return nil, err
		}
		for _, id := range ancestorIDs {
			if _, ok := seen[id]; ok {
				continue
			}
			seen[id] = struct{}{}
			ids = append(ids, id)
		}
	}
	return ac.userRepo.GetUserV1BoundsByIDs(ctx, ids)
}

func (ac *AppUsecase) GetStakingStakedLast(ctx context.Context) (*StakingStaked, error) {
	var (
		rLast *StakingStaked
		err   error
	)
	rLast, err = ac.userRepo.GetStakingStakedLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertStakingStaked(ctx context.Context, trade *StakingStaked) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertStakingStaked(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "bind写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetStakingQueueAddedLast(ctx context.Context) (*StakingQueueAdded, error) {
	var (
		rLast *StakingQueueAdded
		err   error
	)
	rLast, err = ac.userRepo.GetStakingQueueAddedLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertStakingQueueAdded(ctx context.Context, trade *StakingQueueAdded) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertStakingQueueAdded(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "bind写入mysql错误")
		return err
	}

	return err
}
