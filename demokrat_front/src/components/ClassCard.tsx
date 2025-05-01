
import React from 'react';
import { DanceClass } from '@/types';
import { Calendar, Clock, MapPin, DollarSign, User } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { useAuth } from '@/context/AuthContext';
import { useToast } from '@/components/ui/use-toast';
import { useNavigate } from 'react-router-dom';

interface ClassCardProps {
  danceClass: DanceClass;
  onEnroll?: (classId: string) => void;
}

const ClassCard = ({ danceClass, onEnroll }: ClassCardProps) => {
  const { user } = useAuth();
  const { toast } = useToast();
  const navigate = useNavigate();
  
  const isStudent = user?.role === 'student';
  const isTrainer = user?.role === 'trainer';
  const isClassTrainer = user?.id === danceClass.trainer.id;
  
  // Проверяем, записан ли студент на занятие
  const isEnrolled = danceClass.enrolledStudents?.includes(user?.id || '');
  
  const handleEnroll = () => {
    if (!user) {
      toast({
        title: "Требуется авторизация",
        description: "Пожалуйста, войдите в систему для записи на занятия",
        variant: "destructive",
      });
      navigate('/auth');
      return;
    }
    
    if (onEnroll) {
      onEnroll(danceClass.id);
    }
  };

  return (
    <div className="card group">
      <div className="flex flex-col h-full">
        <div className="mb-4 flex items-center justify-between">
          <h3 className="text-lg font-semibold text-white group-hover:text-demokrat-purple transition-colors">
            Танцевальное занятие
          </h3>
          <span className="text-sm text-demokrat-purple font-medium px-2 py-1 bg-demokrat-purple/10 rounded-sm">
            {danceClass.date}
          </span>
        </div>
        
        <div className="space-y-3 flex-grow">
          <div className="flex items-center text-white/70">
            <Clock className="h-4 w-4 mr-2 text-demokrat-purple" />
            <span>{danceClass.time}</span>
          </div>
          
          <div className="flex items-center text-white/70">
            <MapPin className="h-4 w-4 mr-2 text-demokrat-purple" />
            <span>{danceClass.location}</span>
          </div>
          
          <div className="flex items-center text-white/70">
            <DollarSign className="h-4 w-4 mr-2 text-demokrat-purple" />
            <span>{danceClass.price} ₽</span>
          </div>
          
          <div className="flex items-center text-white/70">
            <User className="h-4 w-4 mr-2 text-demokrat-purple" />
            <span>{danceClass.trainer.name}</span>
          </div>
        </div>
        
        <div className="mt-4 pt-4 border-t border-white/10">
          {isStudent && !isEnrolled && (
            <Button 
              onClick={handleEnroll}
              className="w-full bg-demokrat-purple hover:bg-demokrat-purple/90 text-white"
            >
              Записаться
            </Button>
          )}
          
          {isStudent && isEnrolled && (
            <Button 
              disabled
              className="w-full bg-demokrat-purple/50 text-white cursor-not-allowed"
            >
              Вы записаны
            </Button>
          )}
          
          {isTrainer && isClassTrainer && (
            <Button 
              variant="outline"
              className="w-full border-demokrat-purple/30 text-demokrat-purple hover:bg-demokrat-purple/10"
            >
              Редактировать
            </Button>
          )}
          
          {!user && (
            <Button 
              onClick={() => navigate('/auth')}
              className="w-full bg-demokrat-purple hover:bg-demokrat-purple/90 text-white"
            >
              Войдите для записи
            </Button>
          )}
        </div>
      </div>
    </div>
  );
};

export default ClassCard;
