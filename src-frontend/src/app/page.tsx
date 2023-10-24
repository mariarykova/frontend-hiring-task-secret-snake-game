"use client";

import { useCallback, useEffect, useState } from "react";
import SnakeCard from "./components/snake-card";
import { SnakeCardProps, SnakeStage, Bid } from "./types";
import { connectToWebsocket } from "./api/websocket";
import { WinnerSnakeBanner } from "./components/winner-banner";
import { scrollToElement } from "./utils/scrollToElement";
import { handleCardWinner } from "./utils/handleCardWinner";

export default function Home() {
  const [snakesCard, setSnakesCards] = useState<SnakeCardProps[]>([]);
  const [bid, setBid] = useState<Bid>();
  const [isOpenBanner, setIsOpenBanner] = useState(false);
  const [finalBids, setFinalBids] = useState<number[]>([]);
  const [winnerId, setWinnerId] = useState<string>();

  useEffect(() => {
    const fetchSnakes = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/snakes");
        const data = await response.json();
        const newData = data.map((item: any) => ({
          ...item,
          tvl: 0,
          isBidded: false,
        }));
        setSnakesCards(newData);
      } catch (err) {
        console.log(err);
      }
    };
    fetchSnakes();
    connectToWebsocket(setBid);
  }, []);

  useEffect(() => {
    bid && handleBidUpdates(bid);
    setIsOpenBanner(false);
  }, [bid]);

  const handleBidUpdates = (bid: Bid) => {
    switch (bid.stage) {
      case SnakeStage.Stage_One:
        setSnakesCards((prevSnakeData) => {
          const snakeToUpdate = prevSnakeData.find(
            (snake) => snake.id === bid.id
          );

          if (!snakeToUpdate) {
            return [
              ...prevSnakeData,
              {
                id: bid.id,
                tvl: bid.bid,
                isBidded: false,
                bid: bid.bid,
                stage: bid.stage,
              },
            ];
          }

          return prevSnakeData.map((snake) =>
            snake.id === bid.id
              ? {
                  ...snake,
                  tvl: snake.tvl + bid.bid,
                  isBidded: true,
                  bid: bid.bid,
                }
              : { ...snake, isBidded: false }
          );
        });
        scrollToElement(bid.id);
        break;
      case SnakeStage.Stage_Two:
        setSnakesCards((prevSnakeData) =>
          prevSnakeData.map((snake) =>
            snake.id === bid.id
              ? {
                  ...snake,
                  tvl: Math.max(0, snake.tvl - bid.bid),
                  isBidded: true,
                }
              : { ...snake, isBidded: false }
          )
        );
        scrollToElement(bid.id);
        break;
      case SnakeStage.Stage_Three:
        const snakeId = bid.id;
        handleCardWinner({
          snakeId,
          setIsOpenBanner,
          setFinalBids,
          setWinnerId,
        });
        setSnakesCards((prevSnakeData) => {
          const updatedSnakes = prevSnakeData
            .filter((snake) => snake.id !== bid.id)
            .map((snake) => ({
              ...snake,
              isBidded: false,
            }));
          return updatedSnakes;
        });
        break;
      default:
        break;
    }
  };

  const renderCards = useCallback(() => {
    return snakesCard.map((card, index) => {
      return (
        <div key={card.id}>
          <SnakeCard card={card} />
        </div>
      );
    });
  }, [snakesCard]);

  return (
    <main>
      <div className="p-20 max-[600px]:p-0">
        <div className="h-[500px] flex wrap justify-between items-center gap-5 overflow-y-scroll max-[600px]:flex-col h-[100vh]">
          {renderCards()}
        </div>
      </div>
      {isOpenBanner && <WinnerSnakeBanner id={winnerId} data={finalBids} />}
    </main>
  );
}
