import * as fs from "fs";

export function saveDataToFile(jsonData: string) {
  const fileName = "/users.json";
  fs.writeFile("data.json", jsonData, (err: any) => {
    if (err) {
      console.error("Error writing file:", err);
      return;
    }
    console.log("Data saved to data.json");
  });
}
