# 📚 Документация проекта demokrat_front

## 🎯 О проекте

**demokrat_front** — это фронтенд-приложение для управления классами и обучением. Построено на React + TypeScript с использованием Vite как сборщика и Shadcn/UI для компонентов интерфейса.

---

## 📁 Структура папок проекта

```
demokrat_front/
├── src/                           # Основной исходный код
│   ├── pages/                     # Страницы приложения
│   │   ├── Auth.tsx              # Страница авторизации
│   │   ├── Classes.tsx           # Лента классов
│   │   ├── CreateClass.tsx       # Создание нового класса
│   │   ├── Dashboard.tsx         # Личный кабинет пользователя
│   │   ├── Index.tsx             # Главная страница
│   │   └── NotFound.tsx          # Страница 404
│   ├── components/               # Переиспользуемые компоненты
│   │   ├── ClassCard.tsx         # Карточка класса
│   │   ├── ClassForm.tsx         # Форма создания\редакт. класса
│   │   ├── Navbar.tsx            # Верхняя навигационная панель
│   │   └── ui/                   # UI компоненты из Shadcn
│   ├── context/                  # React контексты
│   │   └── AuthContext.tsx       # Контекст аутентификации
│   ├── hooks/                    # Пользовательские React хуки
│   │   ├── use-mobile.tsx        # Хук для определения мобильного устройства
│   │   └── use-toast.ts          # Хук для уведомлений
│   ├── services/                 # API сервисы
│   │   └── api.ts                # Функции для взаимодействия с бэкэндом
│   ├── types/                    # TypeScript типы и интерфейсы
│   │   └── types.ts              # Определение типов данных
│   ├── lib/                      # Утилиты и вспомогательные функции
│   │   └── utils.ts              # Общие утилиты
│   ├── mocks/                    # Тестовые данные (заглушки)
│   │   └── mocks.ts              # Mock данные для разработки
│   ├── assets/                   # Статические ресурсы (картинки, иконки)
│   ├── App.tsx                   # Главный компонент приложения
│   ├── App.css                   # Стили приложения
│   ├── main.tsx                  # Точка входа приложения
│   ├── index.css                 # Глобальные стили
│   └── vite-env.d.ts             # TypeScript декларации Vite
├── public/                       # Статические файлы (не обрабатываются Vite)
├── node_modules/                 # Установленные зависимости
├── .env                          # Переменные окружения (локально)
├── .env.mock                     # Переменные для режима с mock данными
├── package.json                  # Конфигурация проекта и зависимости
├── tsconfig.json                 # Конфигурация TypeScript
├── vite.config.ts                # Конфигурация Vite
├── tailwind.config.ts            # Конфигурация Tailwind CSS
├── postcss.config.js             # Конфигурация PostCSS
├── eslint.config.js              # Конфигурация ESLint
├── components.json               # Конфигурация Shadcn/UI
└── index.html                    # HTML файл приложения
```

---

## 🚀 Быстрый старт

### Установка и запуск

```bash
# Установка зависимостей
npm install

# Запуск dev сервера
npm run dev

# Запуск с mock данными (без подключения к бэкэнду)
npm run dev:mock

# Пока не используется -----------
    # Сборка для production
    npm run build
    # Сборка для development
    npm run build:dev
    # Линтинг кода
    bun lint
    # Предпросмотр собранной версии
    bun preview
# Пока не используется -----------
```

**Приложение будет доступно на:** `http://localhost:8080`

---

## 🛠️ Технологический стек

| Технология | Назначение |
|-----------|----------|
| **React 18** | Библиотека для создания UI |
| **TypeScript** | Типизированный JavaScript |
| **Vite** | Быстрый сборщик и dev сервер |
| **React Router** | Маршрутизация страниц |
| **TanStack Query** | Управление данными и кэширование API запросов |
| **Tailwind CSS** | Утилитарный CSS фреймворк |
| **Shadcn/UI** | Библиотека готовых UI компонентов |
| **React Hook Form** | Управление формами |
| **Zod** | Валидация схем данных |
| **Sonner** | Библиотека для всплывающих уведомлений |
| **Lucide React** | Библиотека иконок |
| **Date-fns** | Работа с датами |

---

## 📊 Архитектура приложения

### Поток данных

```
User Input (Pages/Components)
        ↓
React State / Context (AuthContext)
        ↓
API Service (services/api.ts)
        ↓
Backend (localhost:5088)
        ↓
Database
```

### Principales компоненты

1. **Pages** — Основные страницы приложения, каждая отвечает за один маршрут
2. **Components** — Переиспользуемые компоненты (UI элементы, формы, карточки)
3. **Context** — Глобальное состояние (аутентификация)
4. **Services** — API функции для коммуникации с бэкэндом
5. **Types** — TypeScript типы для всех объектов данных
6. **Hooks** — Пользовательские React хуки для переиспользуемой логики

---

## 🔐 Аутентификация

### AuthContext

Управляет состоянием пользователя и его авторизацией.

**Файл:** `src/context/AuthContext.tsx`

**Функции:**
- `useAuth()` — хук для доступа к контексту аутентификации
- `login(email, password)` — авторизация пользователя
- `register(email, password, name, isTrainer)` — регистрация нового пользователя
- `logout()` — выход из аккаунта

**Использование:**
```typescript
import { useAuth } from '@/context/AuthContext';

export const MyComponent = () => {
  const { user, isAuthenticated, login, logout } = useAuth();
  
  if (!isAuthenticated) return <div>Не авторизованы</div>;
  return <div>Добро пожаловать, {user?.name}!</div>;
};
```

---

## 🌐 API взаимодействие

### API Service

**Файл:** `src/services/api.ts`

Содержит функции для всех API запросов к бэкэнду.

**Ключевые функции:**
- `apiRequest<T>(endpoint, method, data)` — универсальная функция для API запросов
- Автоматическая обработка авторизации с помощью Bearer токена
- Сохранение токенов в `localStorage`

**Пример использования:**
```typescript
import { authService, classesService } from '@/services/api';

// Авторизация
const tokens = await authService.login({ email: 'user@example.com', password: 'pass' });

// Получение классов
const classes = await classesService.getAll();

// Создание класса
await classesService.create({ name: 'Новый класс', date: '2026-01-01' });
```

### Mock режим

Для разработки без бэкэнда используйте mock режим:

```bash
npm run dev:mock
```

Это активирует mock данные из `src/mocks/mocks.ts` вместо реальных API запросов.

**Переменная окружения:** `VITE_USE_MOCK=true` в `.env.mock` 
можно использовать переменную для проверки режима работы приложения (dev c backend или dev:mock с заглушками)

---

## 📝 Типы данных

### Определение типов

**Файл:** `src/types/types.ts`

**Основные типы:**

```typescript
interface User {
  id: string;
  name: string;
  email: string;
  isTrainer: boolean;  // true для тренеров, false для студентов
}

interface Class {
  id: string;
  date: string;        // Формат: YYYY-MM-DD
  time: string;        // Формат: HH:mm
  location: string;
  price: number;
  trainer: {
    id: string;
    name: string;
  };
  enrolledStudents?: string[];  // Массив ID студентов
}

interface Tokens {
  accessToken: string;
  refreshToken: string;
}
```

**Правило:** Всегда используйте эти типы при работе с данными!

---

## 🎨 Компоненты Shadcn/UI

Готовые UI компоненты находятся в `src/components/ui/`

**Примеры использования:**

```typescript
// Кнопка
import { Button } from '@/components/ui/button';
<Button>Нажми меня</Button>

// Форма
import { Form } from '@/components/ui/form';
import { FormField, FormItem, FormLabel, FormControl } from '@/components/ui/form';
// ... использование с react-hook-form

// Карточка
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '@/components/ui/card';
<Card>
  <CardHeader>
    <CardTitle>Заголовок</CardTitle>
  </CardHeader>
  <CardContent>Содержимое</CardContent>
</Card>

// Диалог
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
// ... использование для модальных окон

// Всплывающее уведомление
import { useToast } from '@/components/ui/use-toast';
const { toast } = useToast();
toast({ title: 'Успешно!', description: 'Действие выполнено' });
```

Для добавления новых компонентов Shadcn используйте:
```bash
npx shadcn-ui@latest add [component-name]
```

---

## 📖 Где что находится

### Нужно добавить новую страницу?
1. Создайте файл в `src/pages/YourPage.tsx`
2. Добавьте маршрут в `src/App.tsx` (в компоненте `Routes`)
3. Если нужны компоненты — создавайте их в `src/components/`

### Нужно добавить новый компонент?
- Если это общий компонент, используемый в разных местах → `src/components/`
- Если это компонент для конкретной страницы → создавайте его в той же папке или `src/components/`
- Если это UI элемент, который нужен везде → добавьте в `src/components/ui/`

### Нужно работать с данными с бэкэнда?
1. Добавьте функцию в `src/services/api.ts`
2. Используйте TanStack Query для кэширования: 
   ```typescript
   import { useQuery } from '@tanstack/react-query';
   const { data, isLoading } = useQuery({
     queryKey: ['classes'],
     queryFn: () => classesService.getAll()
   });
   ```

### Нужно добавить новый тип данных?
Добавьте в `src/types/types.ts`:
```typescript
export interface MyNewType {
  id: string;
  name: string;
  // другие поля
}
```

### Нужны новые переменные окружения?
1. Добавьте в `.env` (для обычной разработки)
2. Добавьте в `.env.mock` (для mock режима)
3. Используйте в коде: `import.meta.env.VITE_YOUR_VAR`

---

## 🎯 Основные страницы

| Страница | Файл | Назначение |
|---------|------|----------|
| Главная | `pages/Index.tsx` | Приветственная страница |
| Авторизация | `pages/Auth.tsx` | Вход и регистрация |
| Классы | `pages/Classes.tsx` | Список доступных классов |
| Создать класс | `pages/CreateClass.tsx` | Форма создания класса (только для тренеров) |
| Кабинет | `pages/Dashboard.tsx` | Личный кабинет пользователя |
| 404 | `pages/NotFound.tsx` | Страница не найдена |

---

## 🧩 Основные компоненты

| Компонент | Файл | Назначение |
|-----------|------|----------|
| Navbar | `components/Navbar.tsx` | Навигационная панель в шапке |
| ClassCard | `components/ClassCard.tsx` | Карточка класса (дисплей) |
| ClassForm | `components/ClassForm.tsx` | Форма для создания/редактирования класса |

---

## 🎨 Стили

### Tailwind CSS
Используется утилитарный подход. Классы применяются прямо в JSX:

```typescript
<div className="flex items-center justify-between p-4 bg-white rounded-lg shadow">
  <h2 className="text-lg font-semibold text-gray-800">Заголовок</h2>
  <button className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
    Действие
  </button>
</div>
```

### Глобальные стили
- `src/index.css` — глобальные CSS переменные и стили
- `src/App.css` — стили специфичные для App компонента

### Тема
Цвета и тема настраиваются в `tailwind.config.ts` и `index.css`

---

## 🐛 Debugging

### DevTools

1. **React DevTools** — расширение для браузера для инспектирования компонентов
2. **Redux DevTools** — может быть полезно для отслеживания состояния
3. **Console** — использует `console.log()` для логирования

### Советы

```typescript
// Логирование компонента
export const MyComponent = () => {
  console.log('MyComponent рендерится');
  return <div>Содержимое</div>;
};

// Логирование эффектов
useEffect(() => {
  console.log('Компонент смонтирован');
  return () => console.log('Компонент демонтирован');
}, []);

// Логирование ошибок
try {
  await apiCall();
} catch (error) {
  console.error('Ошибка API:', error);
}
```

---

## 📋 Чек-лист перед коммитом

- [ ] Код проходит линтинг (`bun lint`)
- [ ] Нет ошибок в консоли
- [ ] Все функции правильно типизированы (TypeScript)
- [ ] Используются нужные типы из `types.ts`
- [ ] API функции находятся в `services/api.ts`
- [ ] Переиспользуемые компоненты в `components/`
- [ ] Страницы в `pages/`
- [ ] Нет console.log() в production коде
- [ ] Компоненты правильно структурированы
- [ ] Комментарии добавлены для сложной логики

---

## 🚨 Частые ошибки и решения

### ❌ Импорты от абсолютного пути не работают
**Проблема:** `import { MyComponent } from '../../../components/MyComponent'`

**Решение:** Используйте alias `@`:
```typescript
import { MyComponent } from '@/components/MyComponent'
```

### ❌ Типы не найдены
**Проблема:** TypeScript ошибка про неизвестный тип

**Решение:** Убедитесь, что тип экспортирован из `types.ts`:
```typescript
export interface MyType {
  // ...
}
```

### ❌ API запрос возвращает 401
**Проблема:** Неавторизованный доступ

**Решение:** Проверьте, что:
1. Пользователь авторизован (используйте `useAuth()`)
2. Токен сохранен в `localStorage` с ключом `demokratAccessToken`
3. Бэкэнд запущен на `http://localhost:5088`

### ❌ Mock данные не загружаются
**Проблема:** Приложение пытается подключиться к реальному API

**Решение:**
1. Запустите `bun dev:mock` (не просто `bun dev`)
2. Убедитесь, что `.env.mock` содержит `VITE_USE_MOCK=true`

### ❌ Компонент не переrender'ится при изменении состояния
**Проблема:** UI не обновляется при изменении данных

**Решение:** Проверьте:
1. Используете ли вы `useState()` для управления состоянием
2. Правильно ли указаны dependencies в `useEffect()`
3. Не мутируете ли вы массивы/объекты напрямую

---

## 💡 Лучшие практики

### 1. Всегда используйте типы
```typescript
// ❌ Плохо
const user: any = {};

// ✅ Хорошо
const user: User = { id: '', name: '', email: '', isTrainer: false };
```

### 2. Группируйте импорты
```typescript
// React и библиотеки сверху
import React, { useState } from 'react';
import { useQuery } from '@tanstack/react-query';

// Локальные импорты
import { MyComponent } from '@/components/MyComponent';
import { User } from '@/types/types';
import { userService } from '@/services/api';
```

### 3. Используйте константы для URL и API endpoints
```typescript
const API_BASE_URL = 'http://localhost:5088';
const ENDPOINTS = {
  USERS: '/api/users',
  CLASSES: '/api/classes',
};
```

### 4. Обрабатывайте ошибки
```typescript
const { data, error, isLoading } = useQuery({
  queryKey: ['classes'],
  queryFn: classesService.getAll
});

if (error) return <div>Ошибка загрузки: {error.message}</div>;
if (isLoading) return <div>Загрузка...</div>;
```

### 5. Не забывайте про dependencies в useEffect
```typescript
useEffect(() => {
  fetchData(userId);
}, [userId]); // userId в dependencies!
```

### 6. Используйте React Query для кэширования
```typescript
// Это автоматически кэшируется и переиспользуется
const { data: classes } = useQuery({
  queryKey: ['classes', filterId],
  queryFn: () => classesService.getAll(filterId)
});
```

### 7. Разделяйте логику компонентов
```typescript
// ❌ Не делайте так - весь код в одном компоненте
export const HugeComponent = () => { /* 500 строк кода */ };

// ✅ Делите на части
export const ClassList = () => { /* логика списка */ };
export const ClassItem = (props) => { /* логика элемента */ };
export const ClassForm = () => { /* логика формы */ };
```

---

## 📞 Вопросы и поддержка

При возникновении вопросов:
1. Проверьте эту документацию
2. Посмотрите на примеры в похожих компонентах
3. Проверьте браузерную консоль на ошибки
4. Сдайся слабость! Если не можешь решить сам - обратись к Егору Дмитриевичу

---

## 🔄 Версионирование

**Текущая версия:** 0.0.0

Все обновления документации должны отражать текущее состояние проекта.

---

**Последнее обновление:** 2026-07-22