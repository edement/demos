
import React from 'react';
import { Link } from 'react-router-dom';
import { useAuth } from '@/context/AuthContext';
import Navbar from '@/components/Navbar';
import { Button } from '@/components/ui/button';

const Index = () => {
  const { isAuthenticated, user } = useAuth();

  return (
    <div className="relative min-h-screen flex flex-col">
      <Navbar />
      
      <main className="flex-grow flex flex-col mt-16">
        {/* Hero Section */}
        <section className="flex-grow flex items-center justify-center px-4 py-20 relative overflow-hidden">
          <div className="absolute inset-0 z-0 opacity-20">
            <div className="absolute inset-0 bg-gradient-to-b from-transparent to-demokrat-dark"></div>
            <img 
              src="https://images.unsplash.com/photo-1547153760-18fc86324498?q=80&w=1974&auto=format&fit=crop"
              alt="Dance Studio Background" 
              className="w-full h-full object-cover"
            />
          </div>
          
          <div className="max-w-2xl mx-auto text-center relative z-10 animate-fade-in">
            <h1 className="text-4xl sm:text-5xl md:text-6xl font-bold mb-6">
              <span className="text-white">ДОБРО ПОЖАЛОВАТЬ </span>
              <span className="graffiti-text block sm:inline">DEMOKRAT</span>
            </h1>
            
            <p className="text-lg text-white/80 mb-8 max-w-lg mx-auto">
              Приложение для любителей уличной культуры. 
              Присоединяйтесь к нашему сообществу и выражайте себя в движение.
            </p>
            
            <div className="flex flex-col sm:flex-row gap-4 justify-center">
              {!isAuthenticated ? (
                <Link to="/auth">
                  <Button className="btn-primary min-w-[180px]">
                    Вход / Регистрация
                  </Button>
                </Link>
              ) : (
                <Link to="/dashboard">
                  <Button className="btn-primary min-w-[180px]">
                    Профиль
                  </Button>
                </Link>
              )}
              
              <Link to="/classes">
                <Button variant="outline" className="btn-secondary min-w-[180px]">
                  Посмотреть Занятия
                </Button>
              </Link>
            </div>
          </div>
        </section>
        
        {/* Featured Content */}
        <section className="bg-demokrat-gray py-16 relative">
          <div className="max-w-6xl mx-auto px-4">
            <div className="text-center mb-12">
              <h2 className="text-2xl font-bold text-white mb-2">НАШЕ СООБЩЕСТВО</h2>
              <div className="w-16 h-1 bg-demokrat-purple mx-auto"></div>
            </div>
            
            <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
              <div className="card flex flex-col items-center text-center">
                <div className="w-16 h-16 rounded-full bg-demokrat-purple/20 flex items-center justify-center mb-4">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8 text-demokrat-purple" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                  </svg>
                </div>
                <h3 className="text-xl font-semibold mb-2">Мы не привязаны к месту</h3>
                <p className="text-white/70">Занятия могут проходить где угодно, тебе надо только выбрать!</p>
              </div>
              
              <div className="card flex flex-col items-center text-center">
                <div className="w-16 h-16 rounded-full bg-demokrat-purple/20 flex items-center justify-center mb-4">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8 text-demokrat-purple" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                  </svg>
                </div>
                <h3 className="text-xl font-semibold mb-2">Люди</h3>
                <p className="text-white/70">У нас огромное количество людей обожающих танцы и готовых научить тебя всему.</p>
              </div>
              
              <div className="card flex flex-col items-center text-center">
                <div className="w-16 h-16 rounded-full bg-demokrat-purple/20 flex items-center justify-center mb-4">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8 text-demokrat-purple" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                  </svg>
                </div>
                <h3 className="text-xl font-semibold mb-2">Проверенные Тренера</h3>
                <p className="text-white/70">Все тренера здесь - мировые профессионалы, так что ты точно можешь им доверять.</p>
              </div>
            </div>
          </div>
        </section>
      </main>
      
      <footer className="bg-demokrat-dark py-8 border-t border-white/10">
        <div className="max-w-6xl mx-auto px-4 text-center">
          <p className="text-white/50 text-sm">
            © {new Date().getFullYear()} Demokrat App. All rights reserved.
          </p>
        </div>
      </footer>
    </div>
  );
};

export default Index;
