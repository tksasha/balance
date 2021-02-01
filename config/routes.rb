# frozen_string_literal: true

Rails.application.routes.draw do
  resources :items, only: %i[index create edit update destroy]

  resources :cashes, only: %i[index edit update]

  resources :categories

  resource :backoffice, only: :show, controller: :backoffice

  resources :consolidations, only: :index

  namespace :backoffice do
    resources :exchange_rates, only: :index

    resources :cashes, except: :show
  end

  root to: 'items#index'
end
