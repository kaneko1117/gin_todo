"use client";

import { Checkbox } from "@/components/checkbox";
import { checkedTodoList } from "@/app/action";
import { GraphQLResponse } from "../type";

type Props = {
  data: GraphQLResponse;
};

export const TodoCheckBox = ({ data }: Props) => {
  return (
    <div className="mt-2">
      {data.getTasks.map((item) => {
        return (
          <div key={item.ID} className="flex items-center gap-2">
            <Checkbox
              id={item.ID}
              checked={item.isChecked}
              onCheckedChange={async (isCheck) => {
                if (typeof isCheck !== "boolean") return;
                await checkedTodoList(item.ID, isCheck);
              }}
              className="cursor-pointer"
              name="checked"
            />
            <label htmlFor={item.task} className="text-lg ml-2">
              <p>{item.task}</p>
            </label>
          </div>
        );
      })}
    </div>
  );
};
