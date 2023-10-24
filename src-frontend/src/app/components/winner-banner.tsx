"use client";

import React from "react";
import { WinnerSnakeBannerProps } from "@/app/types";

export const WinnerSnakeBanner = ({ data, id }: WinnerSnakeBannerProps) => {
  const maxBid = Math.max.apply(null, data);
  const filteredBids = data.filter((num) => num !== maxBid);
  const secondBid = Math.max(...filteredBids);

  return (
    <div className="w-[300px] h-[400px] p-5 border-2 border-solid bg-sky-200 border-black text-black absolute top-2/4 left-2/4 transform -translate-x-1/2 -translate-y-1/2">
      <h3 className="pb-4">Bidding finished!!</h3>
      <p className="pb-4">
        Snake Highest bid was{" "}
        <span className="font-bold">{maxBid > Infinity ? maxBid : 0}</span>
      </p>
      <p className="pb-4">
        The Second to Highest bid was{" "}
        <span className="font-bold">
          {secondBid > Infinity ? secondBid : 0}
        </span>
      </p>
      <h4>
        Congradulations to the lucky winner of{" "}
        <span className="font-bold">{id}</span>
      </h4>
    </div>
  );
};
