package elo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type EloSuite struct {
	suite.Suite
}

func (suite *EloSuite) TestProbability() {
	ranking1 := 1000.0
	ranking2 := 1200.0

	suite.Equal(Probability(ranking1, ranking2), 0.24)
	suite.Equal(Probability(ranking2, ranking1), 0.76)
}

func (suite *EloSuite) TestRankingWithBetterPlayerWinning() {
	ranking1 := 1200.0
	ranking2 := 1000.0
	K := 30

	new1, new2 := Ranking(ranking1, ranking2, true, K)

	suite.Equal(new1, 1207.2)
	suite.Equal(new2, 992.8)
}

func (suite *EloSuite) TestRankingWithWorsePlayerWinning() {
	ranking1 := 1200.0
	ranking2 := 1000.0
	K := 30

	new1, new2 := Ranking(ranking1, ranking2, false, K)

	suite.Equal(new1, 1177.2)
	suite.Equal(new2, 1022.8)
}

func TestEloSuite(t *testing.T) {
	suite.Run(t, new(EloSuite))
}
