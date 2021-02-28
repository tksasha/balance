# frozen_string_literal: true

Rails.application.routes.draw do
  scope '/(:currency)' do
    resources :items, only: %i[create edit update destroy]

    resources :cashes, only: %i[index edit update]

    resources :categories

    resource :backoffice, only: :show, controller: :backoffice

    resources :consolidations, only: :index

    namespace :backoffice do
      resources :exchange_rates, only: :index

      resources :cashes, except: :show
    end

    get '(/:month)(/:category_id)', to: 'items#index', as: :root
  end
end
