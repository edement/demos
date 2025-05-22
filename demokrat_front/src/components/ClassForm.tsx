
import React from 'react';
import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { useToast } from '@/components/ui/use-toast';
import { useAuth } from '@/context/AuthContext';
import { Calendar, Clock, MapPin, DollarSign } from 'lucide-react';

interface ClassFormData {
  date: string;
  time: string;
  location: string;
  price: number;
}

interface ClassFormProps {
  onSubmit: (data: ClassFormData) => void;
  isLoading?: boolean;
}

const ClassForm = ({ onSubmit, isLoading = false }: ClassFormProps) => {
  const { register, handleSubmit, formState: { errors } } = useForm<ClassFormData>();
  const { toast } = useToast();
  const { user } = useAuth();

  const submitHandler = (data: ClassFormData) => {
    if (!user.isTrainer) {
      toast({
        title: "Permission denied",
        description: "Only trainers can create classes",
        variant: "destructive",
      });
      return;
    }

    onSubmit(data);
  };

  return (
    <form onSubmit={handleSubmit(submitHandler)} className="space-y-6">
      <div className="space-y-2">
        <label htmlFor="date" className="flex items-center text-sm font-medium text-white">
          <Calendar className="h-4 w-4 mr-2 text-demokrat-purple" />
          Date
        </label>
        <input
          id="date"
          type="date"
          className="input-field w-full"
          {...register("date", { required: "Date is required" })}
        />
        {errors.date && (
          <p className="text-red-500 text-xs mt-1">{errors.date.message}</p>
        )}
      </div>
      
      <div className="space-y-2">
        <label htmlFor="time" className="flex items-center text-sm font-medium text-white">
          <Clock className="h-4 w-4 mr-2 text-demokrat-purple" />
          Time
        </label>
        <input
          id="time"
          type="time"
          className="input-field w-full"
          {...register("time", { required: "Time is required" })}
        />
        {errors.time && (
          <p className="text-red-500 text-xs mt-1">{errors.time.message}</p>
        )}
      </div>
      
      <div className="space-y-2">
        <label htmlFor="location" className="flex items-center text-sm font-medium text-white">
          <MapPin className="h-4 w-4 mr-2 text-demokrat-purple" />
          Location
        </label>
        <input
          id="location"
          type="text"
          className="input-field w-full"
          placeholder="Room 101, Studio Demokrat"
          {...register("location", { required: "Location is required" })}
        />
        {errors.location && (
          <p className="text-red-500 text-xs mt-1">{errors.location.message}</p>
        )}
      </div>
      
      <div className="space-y-2">
        <label htmlFor="price" className="flex items-center text-sm font-medium text-white">
          <DollarSign className="h-4 w-4 mr-2 text-demokrat-purple" />
          Price (RUB)
        </label>
        <input
          id="price"
          type="number"
          min="0"
          step="100"
          className="input-field w-full"
          placeholder="1000"
          {...register("price", { 
            required: "Price is required",
            min: { value: 0, message: "Price must be a positive number" },
            valueAsNumber: true
          })}
        />
        {errors.price && (
          <p className="text-red-500 text-xs mt-1">{errors.price.message}</p>
        )}
      </div>
      
      <div className="pt-2">
        <Button 
          type="submit" 
          className="w-full bg-demokrat-purple hover:bg-demokrat-purple/90"
          disabled={isLoading}
        >
          {isLoading ? "Creating..." : "Create Class"}
        </Button>
      </div>
    </form>
  );
};

export default ClassForm;
