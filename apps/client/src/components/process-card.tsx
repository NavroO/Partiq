import type { Process } from "../types/types";

interface ProcessCardProps {
  process: Process
}

export const ProcessCard = ({ process }: ProcessCardProps) => {
  return (
    <div className="transition-all duration-300 hover:shadow-md">
      <div className="pb-2">
        <div className="flex justify-between items-start">

        </div>
        <p className="text-xl mt-2">{process.title}</p>
      </div>
      <div>
        <p className="text-gray-600 mb-4 line-clamp-3">{process.description}</p>
        <div className="flex items-center text-sm text-gray-500 mb-1">
          <span>
            {process.start_at} - {process.end_at}
          </span>
        </div>
      </div>
    </div>
  )
}
