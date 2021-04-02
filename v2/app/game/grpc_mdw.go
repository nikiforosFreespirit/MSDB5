package game

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/dom/team"
	"github.com/mcaci/msdb5/v2/pb"
	"google.golang.org/grpc"
)

func ScoreGrpc(g *Game) string {
	t1, t2 := briscola5.ToGeneralPlayers(g.players).Part(briscola5.IsInCallers(&g.players))

	conn := getConn()
	defer conn.Close()
	client := pb.NewBriscolaClient(conn)

	toPBCards := func(cards set.Cards) *pb.Cards {
		pbcards := make([]*pb.CardID, len(cards))
		for i := range pbcards {
			pbcards[i] = &pb.CardID{Id: uint32(cards[i].ToID())}
		}
		return &pb.Cards{Cards: pbcards}
	}

	s1, err := client.Score(context.Background(), toPBCards(team.CommonPile(t1)))
	if err != nil {
		log.Println(err)
	}
	log.Println(s1.GetPoints())

	s2, err := client.Score(context.Background(), toPBCards(team.CommonPile(t2)))
	if err != nil {
		log.Println(err)
	}
	log.Println(s2.GetPoints())

	return fmt.Sprintf("[%s: %d], [%s: %d]", "Caller team", s1.Points, "Non Caller team", s2.Points)
}

func getConn() *grpc.ClientConn {
	const serverAddr = "localhost:8081"
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Println("error found", err)
	}
	return conn
}
