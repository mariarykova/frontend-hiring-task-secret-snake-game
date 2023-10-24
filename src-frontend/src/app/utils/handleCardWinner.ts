import { HandleCardWinner } from "../types";

export async function handleCardWinner({
  snakeId,
  setIsOpenBanner,
  setFinalBids,
  setWinnerId,
}: HandleCardWinner) {
  try {
    const response = await fetch(
      `http://localhost:8080/api/bids?snake-id=${snakeId}`
    );

    if (!response.ok) {
      throw new Error("Network response was not ok");
    }

    const data = await response.json();

    if (data) {
      setIsOpenBanner(true);
      setFinalBids(data);
      setWinnerId(snakeId);
    }
  } catch (error) {
    console.error("Error fetching data:", error);
  }
}
