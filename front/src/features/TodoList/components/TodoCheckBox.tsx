"use client";

import { Checkbox } from "@/components/checkbox";
import { checkedTodoList } from "@/app/action";
import { todoData } from "../type";

type Props = {
  data: todoData;
};

export const TodoCheckBox = ({ data }: Props) => {
  return (
    <div className="mt-2">
      {data.map((item) => {
        return (
          <div key={item.id} className="flex items-center gap-2">
            <Checkbox
              id={item.task}
              checked={item.checked}
              onCheckedChange={async (isCheck) => {
                if (typeof isCheck !== "boolean") return;
                await checkedTodoList(item.id, isCheck);
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
