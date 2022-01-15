// Code generated by sqlc. DO NOT EDIT.

package wordle

import (
	"context"
)

type Querier interface {
	CountAccountsByDiscordId(ctx context.Context, discordID string) (int64, error)
	CountNicknameByDiscordIdAndServerId(ctx context.Context, arg CountNicknameByDiscordIdAndServerIdParams) (int64, error)
	CountScoresByDiscordId(ctx context.Context, discordID string) (int64, error)
	CreateAccount(ctx context.Context, discordID string) (Account, error)
	CreateNickname(ctx context.Context, arg CreateNicknameParams) (Nickname, error)
	CreateQuipForScore(ctx context.Context, arg CreateQuipForScoreParams) (Quip, error)
	CreateScore(ctx context.Context, arg CreateScoreParams) (WordleScore, error)
	DeleteAccount(ctx context.Context, discordID string) error
	DeleteNickname(ctx context.Context, discordID string) error
	DeleteScoresForUser(ctx context.Context, discordID string) error
	GetAccount(ctx context.Context, discordID string) (Account, error)
	GetNickname(ctx context.Context, arg GetNicknameParams) (Nickname, error)
	GetNicknamesByDiscordId(ctx context.Context, discordID string) ([]Nickname, error)
	GetQuipByScore(ctx context.Context, arg GetQuipByScoreParams) (Quip, error)
	GetQuipsByCreatedByAccount(ctx context.Context, createdByAccount string) ([]Quip, error)
	GetScoreHistoryByAccount(ctx context.Context, arg GetScoreHistoryByAccountParams) ([]GetScoreHistoryByAccountRow, error)
	GetScoresByServerId(ctx context.Context, serverID string) ([]GetScoresByServerIdRow, error)
	GetScoresByServerIdLastWeek(ctx context.Context, serverID string) ([]GetScoresByServerIdLastWeekRow, error)
	ListAccounts(ctx context.Context) ([]Account, error)
	ListNicknames(ctx context.Context) ([]Nickname, error)
	ListScores(ctx context.Context) ([]WordleScore, error)
	UpdateNickname(ctx context.Context, arg UpdateNicknameParams) (Nickname, error)
	UpdateScore(ctx context.Context, arg UpdateScoreParams) (WordleScore, error)
	UpdateTimeZone(ctx context.Context, arg UpdateTimeZoneParams) (Account, error)
}

var _ Querier = (*Queries)(nil)
