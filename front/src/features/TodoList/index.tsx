"use client";
import { Button } from "@/components/button";
import { Checkbox } from "@/components/checkbox";
import { Input } from "@/components/input";
import { todoData } from "./type";
import { useState } from "react";

type Props = {
  data: todoData;
};

export const TodoList = ({ data }: Props) => {
  const [checked, setChecked] = useState<boolean>(false);
  return (
    <div className="flex flex-col items-center gap-5 h-screen mt-5">
      <div className="w-[500px]">
        <form className="flex gap-2">
          <Input type="text" placeholder="タスクを追加" id="todo" name="todo" />
          <Button type="submit" className="cursor-pointer">
            追加
          </Button>
        </form>
        <div className="mt-5">
          {data.map((item) => {
            return (
              <div key={item.id} className="flex items-center gap-2">
                <Checkbox
                  id={item.task}
                  checked={checked}
                  onClick={() => setChecked(!checked)}
                  className="cursor-pointer"
                />
                <label htmlFor={item.task} className="text-lg ml-2">
                  <p>{item.task}</p>
                </label>
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );
};
