import { useState, useEffect } from "react";
import type { Process } from "../types/types";
import { ProcessCard } from "../components/process-card";

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
        <ProcessCard key={p.id} process={p} />
      ))}
    </div>
  );
}
