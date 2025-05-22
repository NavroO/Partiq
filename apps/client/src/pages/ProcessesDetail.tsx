import { useState, useEffect } from "react";
import type { Process } from "../types/types";
import { ProcessCard } from "../components/process-card";
import { useParams } from "react-router-dom";

export default function ProcessesDetail() {
  const [processes, setProcesses] = useState<Process | null>(null);
  const { processID } = useParams();

  useEffect(() => {
    fetch(`http://localhost:8080/processes/${processID}`)
      .then((res) => res.json())
      .then((data) => setProcesses(data))
      .catch((err) => console.error("❌ Błąd pobierania procesów:", err));
  }, [processID]);

  return (
    <div className="p-6 max-w-2xl mx-auto">
      {processes && <ProcessCard process={processes} />}
    </div>
  );
}
