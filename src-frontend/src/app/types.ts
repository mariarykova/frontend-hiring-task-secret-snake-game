export interface SnakeCardProps {
  bid: number;
  id: string;
  stage: number;
  tvl: number;
  isBidded: boolean;
}

export enum SnakeStage {
  Stage_One = 1,
  Stage_Two = 2,
  Stage_Three = 3,
}

export interface WinnerSnakeBannerProps {
  maxBid?: number;
  secondBid?: number;
  id?: string;
  data: number[];
}

export interface Bid {
  id: string;
  bid: number;
  stage: number;
}

export interface HandleCardWinner {
  snakeId: string;
  setIsOpenBanner: (value: boolean) => void;
  setFinalBids: (data: number[]) => void;
  setWinnerId: (value: string) => void;
}
