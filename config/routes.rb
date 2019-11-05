# frozen_string_literal: true

Rails.application.routes.draw do
  resources :items, only: %i[index create edit update destroy]

  resources :cashes, :categories

  resource :backoffice, only: :show, controller: :backoffice

  resources :consolidations, only: :index

  resources :currencies, only: :index

  root to: 'items#index'
end
