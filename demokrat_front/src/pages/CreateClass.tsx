
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Navbar from '@/components/Navbar';
import ClassForm from '@/components/ClassForm';
import { useAuth } from '@/context/AuthContext';
import { useToast } from '@/components/ui/use-toast';
import { classesService } from '@/services/api';

interface ClassFormData {
  date: string;
  time: string;
  location: string;
  price: number;
}

const CreateClass = () => {
  const [isSubmitting, setIsSubmitting] = useState(false);
  const navigate = useNavigate();
  const { user } = useAuth();
  const { toast } = useToast();
  
  // Redirect if not a trainer
  React.useEffect(() => {
    if (user && user.role !== 'trainer') {
      toast({
        title: "Доступ запрещен",
        description: "Только тренеры могут создавать занятия",
        variant: "destructive",
      });
      navigate('/classes');
    }
  }, [user, navigate, toast]);

  const handleCreateClass = async (data: ClassFormData) => {
    setIsSubmitting(true);
    
    try {
      await classesService.createClass(data);
      
      toast({
        title: "Занятие создано",
        description: `Ваше занятие на ${data.date} в ${data.time} было успешно создано.`,
      });
      
      navigate('/classes');
    } catch (error) {
      console.error('Error creating class:', error);
      toast({
        title: "Ошибка",
        description: "Не удалось создать занятие. Пожалуйста, попробуйте снова.",
        variant: "destructive",
      });
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="min-h-screen flex flex-col">
      <Navbar />
      
      <main className="flex-grow p-4 mt-16 max-w-xl mx-auto w-full">
        <div className="mb-8">
          <h1 className="text-2xl font-bold text-white">Создать новое занятие</h1>
          <p className="text-white/70 mt-1">Заполните данные для создания нового танцевального занятия.</p>
        </div>
        
        <div className="card">
          <ClassForm onSubmit={handleCreateClass} isLoading={isSubmitting} />
        </div>
      </main>
    </div>
  );
};

export default CreateClass;
