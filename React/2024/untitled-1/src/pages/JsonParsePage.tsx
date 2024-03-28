import { useEffect, useState } from "react";
import { saveDataToFile } from "./jsonSave";

interface User {
  id: number;
  name: string;
  email: string;
}

export function JsonParsePage() {
  const [users, setUsers] = useState<User[]>([]);
  const [newId, setNewId] = useState(0);
  const [newUserName, setNewUserName] = useState<string>("");
  const [newUserEmail, setNewUserEmail] = useState<string>("");
  const [isReloading, setReload] = useState(true);

  useEffect(() => {
    fetch("/users.json")
      .then((response) => response.json())
      .then((data: User[]) => setUsers(data))
      .catch((error) => console.error("Error fetching users", error));
    setReload(false);
  }, [isReloading]);

  const handleAddUser = () => {
    const newUser: User = {
      id: users.length + 1, // Генерируем новый ID
      name: newUserName,
      email: newUserEmail,
    };
    setUsers([...users, newUser]);
    setNewUserName("");
    setNewUserEmail("");
    setNewId(0);
  };

  const handleUpdateUser = () => {
    for (var i = 0; i < users.length; i++) {
      if (users[i].id === newId) {
        users[i].name = newUserName;
        users[i].email = newUserEmail;
      }
    }
    setUsers(users);
    setNewUserName("");
    setNewUserEmail("");
    setNewId(0);
  };

  const handleSave = () => {
    const usersJson = JSON.stringify(users);
    console.log(usersJson);
    saveDataToFile(usersJson);
    setReload(true);
  };

  const handleCancel = () => {
    setReload(true);
  };

  return (
    <div className="flex justify-center pt-8">
      <div className="w-2/3 max-w-md flex flex-col justify-center">
        <h1 className="text-4xl font-black text-gray-900 text-center">
          Users list
        </h1>
        <ul className="pt-5">
          {users.map((user) => (
            <li key={user.id}>
              <strong>
                {user.id}. {user.name}
              </strong>{" "}
              - {user.email}
            </li>
          ))}
        </ul>
        <input
          type="number"
          className="border-2 rounded border-black"
          placeholder="Name"
          value={newId}
          onChange={(e) => setNewId(Number(e.target.value))}
        />
        <input
          type="text"
          className="border-2 rounded border-black"
          placeholder="Name"
          value={newUserName}
          onChange={(e) => setNewUserName(e.target.value)}
        />
        <input
          type="email"
          placeholder="Email"
          className="border-2 rounded border-black"
          value={newUserEmail}
          onChange={(e) => setNewUserEmail(e.target.value)}
        />
        <button onClick={handleAddUser}>Add User</button>
        <button onClick={handleUpdateUser}>Update User</button>
        <button onClick={handleSave}>Save</button>
        <button onClick={handleCancel}>Cancel</button>
      </div>
    </div>
  );
}
