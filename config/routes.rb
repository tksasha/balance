# frozen_string_literal: true

Rails.application.routes.draw do
  resources :items, only: %i[index create edit update destroy]

  resources :cashes, :categories

  resource :backoffice, only: :show, controller: :backoffice

  resources :consolidations, only: :index

  namespace :backoffice do
    resources :exchange_rates, only: :index
  end

  root to: 'items#index'
end
