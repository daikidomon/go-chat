from django.urls import path

from . import views

app_name = 'chatroom'

urlpatterns = [
    path('', views.IndexView.as_view(), name="index"),
    path('create/', views.CreateRoomView.as_view(), name="create"),
    path('detail/<uuid:pk>', views.DetailRoomView.as_view(), name="detail"),
    path('list/', views.ListRoomView.as_view(), name="list"),
]