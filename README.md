# hse_arch
<b>Тема:</b> Разработка программного модуля для настройки доступа к Государственному информационному ресурсу бухгалтерской (финансовой) отчетности по интерфейсу REST API

Заказчику необходимо автоматизировать процесс анализа финансовых отчетностей, сейчас пользователю требуется скачать отчет и вручную вводить бухгалтерский баланс для анализа.

<b>Далее в тексте:</b>
<br/>ГИР БФО – Государственный информационный ресурс бухгалтерской (финансовой) отчетности

<br/>
<ul>
 <li><b>Пользователи:</b> бухгалтеры, финансовые аналитики, инвесторы</li>
  <li><b>Требования:</b>
    <ul>
      <li>Пользователь должен иметь возможность искать финансовую отчетность компании по различным параметрам</li>
      <li>Пользователь должен получать основную информацию о финансах компании</li>
      <li>Пользователь должен получать результат анализа финансовых данных на предмет экономической деятельности, платежеспособности и технологичности контрагентов</li>
      <li>Пользователь должен иметь возможность экспортировать финансовую отчетность и результаты анализа</li>
      <li>Пользователь должен иметь возможность сохранять компанию в избранное, для быстрого доступа</li>
      <li>Пользователь должен иметь доступ к истории своих запросов</li>
      <li>Система должна обрабатывать информацию полученную от API ГИР БФО и преобразовывать в формат, который требуется для проведения анализа</li>
    </ul>
  </li>
 <li><b>Дополнительный контекст:</b>
   <ul>
      <li>Система основывается на работе с API ГИР БФО</li>
     <li>Анализ производится отдельными сервисами, доступ к которым производится посредством API</li>    
    </ul>
 </li>
</ul>
