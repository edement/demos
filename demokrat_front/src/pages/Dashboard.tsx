
import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Navbar from '@/components/Navbar';
import { useAuth } from '@/context/AuthContext';
import { Button } from '@/components/ui/button';
import { DanceClass } from '@/types';
import ClassCard from '@/components/ClassCard';
import { Link } from 'react-router-dom';
import { PlusCircle, UserCircle } from 'lucide-react';
import { classesService, enrollmentService } from '@/services/api';
import { useToast } from '@/components/ui/use-toast';

const Dashboard = () => {
  const { user, isAuthenticated } = useAuth();
  const navigate = useNavigate();
  const [userClasses, setUserClasses] = useState<DanceClass[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const { toast } = useToast();
  
  const isTrainer = user?.role === 'trainer';

  // Redirect if not authenticated
  useEffect(() => {
    if (!isAuthenticated) {
      navigate('/auth');
    }
  }, [isAuthenticated, navigate]);

  useEffect(() => {
    if (!user) return;
    
    const fetchUserClasses = async () => {
      setIsLoading(true);
      try {
        let classesData: DanceClass[];
        
        if (isTrainer) {
          // Получаем классы тренера
          classesData = await classesService.getTrainerClasses();
        } else {
          // Получаем классы, на которые записан студент
          classesData = await enrollmentService.getUserEnrollments();
        }
        
        setUserClasses(classesData);
      } catch (error) {
        console.error('Error fetching user classes:', error);
        toast({
          title: "Ошибка загрузки",
          description: "Не удалось загрузить ваши занятия. Пожалуйста, попробуйте позже.",
          variant: "destructive",
        });
      } finally {
        setIsLoading(false);
      }
    };
    
    fetchUserClasses();
  }, [user, isTrainer, toast]);

  const handleEnroll = async (classId: string) => {
    try {
      await enrollmentService.enrollToClass(classId);
      toast({
        title: "Запись выполнена",
        description: "Вы успешно записались на занятие",
      });
      
      // Обновляем список занятий
      if (!isTrainer) {
        const updatedClasses = await enrollmentService.getUserEnrollments();
        setUserClasses(updatedClasses);
      }
    } catch (error) {
      console.error(`Error enrolling in class: ${classId}`, error);
      toast({
        title: "Ошибка записи",
        description: "Не удалось записаться на занятие. Пожалуйста, попробуйте снова.",
        variant: "destructive",
      });
    }
  };

  if (!user) {
    return null; // or a loading indicator
  }

  return (
    <div className="min-h-screen flex flex-col">
      <Navbar />
      
      <main className="flex-grow p-4 mt-16 max-w-7xl mx-auto w-full">
        <div className="mb-8">
          <div className="flex items-center gap-4 mb-4">
            <div className="w-16 h-16 rounded-full bg-demokrat-purple/20 flex items-center justify-center">
              <UserCircle className="w-10 h-10 text-demokrat-purple" />
            </div>
            <div>
              <h1 className="text-2xl font-bold text-white">{user.name}</h1>
              <p className="text-white/70">
                {isTrainer ? 'Тренер' : 'Ученик'} • {user.email}
              </p>
            </div>
          </div>
          
          {isTrainer && (
            <div className="flex justify-end">
              <Link to="/create-class">
                <Button className="bg-demokrat-purple hover:bg-demokrat-purple/90">
                  <PlusCircle className="mr-2 h-4 w-4" />
                  Создать новое занятие
                </Button>
              </Link>
            </div>
          )}
        </div>
        
        <div className="mb-8">
          <h2 className="text-xl font-semibold text-white mb-4">
            {isTrainer ? 'Ваши занятия' : 'Записи на занятия'}
          </h2>
          
          {isLoading ? (
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 animate-pulse">
              {[1, 2].map((i) => (
                <div key={i} className="bg-demokrat-gray h-64 rounded-md border border-white/5"></div>
              ))}
            </div>
          ) : (
            <>
              {userClasses.length === 0 ? (
                <div className="card p-8 text-center">
                  <p className="text-white/70 mb-4">
                    {isTrainer 
                      ? "Вы еще не создали ни одного занятия." 
                      : "Вы еще не записаны ни на одно занятие."}
                  </p>
                  
                  {isTrainer ? (
                    <Link to="/create-class">
                      <Button className="bg-demokrat-purple hover:bg-demokrat-purple/90">
                        Создать первое занятие
                      </Button>
                    </Link>
                  ) : (
                    <Link to="/classes">
                      <Button className="bg-demokrat-purple hover:bg-demokrat-purple/90">
                        Просмотреть доступные занятия
                      </Button>
                    </Link>
                  )}
                </div>
              ) : (
                <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
                  {userClasses.map((danceClass) => (
                    <ClassCard key={danceClass.id} danceClass={danceClass} onEnroll={isTrainer ? undefined : handleEnroll} />
                  ))}
                </div>
              )}
            </>
          )}
        </div>
        
        {isTrainer && (
          <div className="card p-6">
            <h3 className="text-lg font-semibold text-white mb-4">Быстрая статистика</h3>
            <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
              <div className="p-4 bg-demokrat-purple/10 rounded-md border border-demokrat-purple/20">
                <p className="text-sm text-white/70 mb-1">Всего занятий</p>
                <p className="text-2xl font-bold text-white">{userClasses.length}</p>
              </div>
              <div className="p-4 bg-demokrat-purple/10 rounded-md border border-demokrat-purple/20">
                <p className="text-sm text-white/70 mb-1">Всего учеников</p>
                <p className="text-2xl font-bold text-white">
                  {userClasses.reduce((total, cls) => total + (cls.enrolledStudents?.length || 0), 0)}
                </p>
              </div>
              <div className="p-4 bg-demokrat-purple/10 rounded-md border border-demokrat-purple/20">
                <p className="text-sm text-white/70 mb-1">Предстоящие занятия</p>
                <p className="text-2xl font-bold text-white">{userClasses.length}</p>
              </div>
            </div>
          </div>
        )}
      </main>
    </div>
  );
};

export default Dashboard;
