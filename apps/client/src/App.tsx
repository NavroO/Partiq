import { BrowserRouter, Route, Routes } from "react-router-dom";
import Home from "./pages/Home";
import ProcessesDetail from "./pages/ProcessesDetail";

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/process/:processID" element={<ProcessesDetail />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
