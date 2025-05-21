import { useState, useEffect } from "react";

type Process = {
  id: number;
  title: string;
  description: string;
  start_at: string;
  end_at: string;
};

export default function Home() {
  const [processes, setProcesses] = useState<Process[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/processes")
      .then((res) => res.json())
      .then((data) => setProcesses(data))
      .catch((err) => console.error("❌ Błąd pobierania procesów:", err));
  }, []);

  return (
    <div className="p-6 max-w-2xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">Dostępne procesy</h1>
      {processes.map((p) => (
        <div key={p.id} className="border rounded-lg p-4 shadow mb-4">
          <h2 className="text-xl font-semibold">{p.title}</h2>
          <p className="text-gray-600 mb-2">{p.description}</p>
          <p className="text-sm text-gray-500">
            {new Date(p.start_at).toLocaleDateString()} –{" "}
            {new Date(p.end_at).toLocaleDateString()}
          </p>
        </div>
      ))}
    </div>
  );
}
