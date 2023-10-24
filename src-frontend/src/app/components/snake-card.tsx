import { SnakeCardProps } from "../types";
import Image from "next/image";

const SnakeCard = ({ card }: { card: SnakeCardProps }) => {
  return (
    <div
      key={card.id}
      id={card.id}
      className={`
      w-[250px] 
      h-[300px] 
      border-2 
      border-solid 
      bg-sky-200 
      border-black 
      text-black 
      p-5
      transform transition duration-300 ease-in-out
    ${card.isBidded ? "scale-110" : "scale-100"}`}
    >
      <Image
        src="/snake.jpeg"
        width={200}
        height={100}
        alt="Picture of the snake"
      />
      <h3>
        Snake ID <span className="font-bold">{card.id}</span>
      </h3>
      <div>
        Snake TVL <span className="font-bold">{card.tvl}</span>
      </div>
      <div>
        Highest bid <span className="font-bold">{card.bid}</span>
      </div>
    </div>
  );
};

export default SnakeCard;
