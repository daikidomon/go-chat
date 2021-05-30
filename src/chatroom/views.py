from django.shortcuts import render
from django.views.generic import TemplateView, CreateView, DetailView, ListView
from django.contrib.auth.decorators import login_required
from django.utils.decorators import method_decorator

from .models import Room

@method_decorator(login_required, name='dispatch')
class IndexView(TemplateView):
    template_name = 'chatroom/index.html'


@method_decorator(login_required, name='dispatch')
class CreateRoomView(CreateView):
    model = Room
    fields = [
        'name',
    ]
    template_name = 'chatroom/form/create.html'


@method_decorator(login_required, name='dispatch')
class DetailRoomView(DetailView):
    model = Room
    template_name = 'chatroom/detail.html'

@method_decorator(login_required, name='dispatch')
class ListRoomView(ListView):
    model = Room
    template_name = 'chatroom/list.html'
