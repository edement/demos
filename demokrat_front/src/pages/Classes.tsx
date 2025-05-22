
import { useState, useEffect } from 'react';
import Navbar from '@/components/Navbar';
import ClassCard from '@/components/ClassCard';
import { Class } from '@/types';
import { useAuth } from '@/context/AuthContext';
import { Button } from '@/components/ui/button';
import { Link } from 'react-router-dom';
import { PlusCircle } from 'lucide-react';
import { classesService, enrollmentService } from '@/services/api';
import { useToast } from '@/components/ui/use-toast';

const Classes = () => {
  const [classes, setClasses] = useState<Class[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const { user } = useAuth();
  const { toast } = useToast();
  
  const isTrainer = user.isTrainer;

  useEffect(() => {
    fetchClasses();
  }, []);

  const fetchClasses = async () => {
    setIsLoading(true);
    try {
      const classesData = await classesService.getAllClasses();
      setClasses(classesData);
    } catch (error) {
      console.error('Error fetching classes:', error);
      toast({
        title: "Ошибка загрузки",
        description: "Не удалось загрузить список занятий. Пожалуйста, попробуйте позже.",
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  };

  const handleEnroll = async (classId: string) => {
    if (!user) return;
    
    try {
      await enrollmentService.enrollToClass(classId);
      toast({
        title: "Запись выполнена",
        description: "Вы успешно записались на занятие",
      });
      // Обновляем список занятий
      fetchClasses();
    } catch (error) {
      console.error(`Error enrolling in class: ${classId}`, error);
      toast({
        title: "Ошибка записи",
        description: "Не удалось записаться на занятие. Пожалуйста, попробуйте снова.",
        variant: "destructive",
      });
    }
  };

  return (
    <div className="min-h-screen flex flex-col">
      <Navbar />
      
      <main className="flex-grow p-4 mt-16 max-w-7xl mx-auto w-full">
        <div className="flex items-center justify-between mb-8">
          <h1 className="text-2xl font-bold text-white">Занятия</h1>
          
          {isTrainer && (
            <Link to="/create-class">
              <Button className="bg-demokrat-purple hover:bg-demokrat-purple/90">
                <PlusCircle className="mr-2 h-4 w-4" />
                Создать занятие
              </Button>
            </Link>
          )}
        </div>
        
        {isLoading ? (
          <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 animate-pulse">
            {[1, 2, 3, 4].map((i) => (
              <div key={i} className="bg-demokrat-gray h-64 rounded-md border border-white/5"></div>
            ))}
          </div>
        ) : (
          <>
            {classes.length === 0 ? (
              <div className="text-center py-16">
                <p className="text-white/70 mb-4">В данный момент нет доступных занятий.</p>
                {isTrainer && (
                  <Link to="/create-class">
                    <Button className="bg-demokrat-purple hover:bg-demokrat-purple/90">
                      Создать первое занятие
                    </Button>
                  </Link>
                )}
              </div>
            ) : (
              <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
                {classes.map((danceClass) => (
                  <ClassCard 
                    key={danceClass.id}
                    danceClass={danceClass}
                    isEnrolled={danceClass.enrolledStudents?.includes(user?.id || '')}
                    onEnroll={handleEnroll}
                  />
                ))}
              </div>
            )}
          </>
        )}
      </main>
    </div>
  );
};

export default Classes;
