import { Toaster } from "@/components/ui/toaster";
import { Toaster as Sonner } from "@/components/ui/sonner";
import { TooltipProvider } from "@/components/ui/tooltip";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { AuthProvider } from "@/context/AuthContext";
import Index from "./pages/Index";
import Classes from "./pages/Classes";
import Auth from "./pages/Auth";
import Dashboard from "./pages/Dashboard";
import CreateClass from "./pages/CreateClass";
import NotFound from "./pages/NotFound";

const queryClient = new QueryClient();

const App = () => (
  <QueryClientProvider client={queryClient}>
    <TooltipProvider>
      <AuthProvider>
        <Toaster/>
        <Sonner/>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Index/>}/>
            <Route path="/classes"element={<Classes/>}/>
            <Route path="/auth"element={<Auth/>}/>
            <Route path="/dashboard" element={<Dashboard />}/>
            <Route path="/create-class" element={<CreateClass />}/>
            <Route path="*" element={<NotFound />}/>
          </Routes>
        </BrowserRouter>
      </AuthProvider>
    </TooltipProvider>
  </QueryClientProvider>
);

export default App;