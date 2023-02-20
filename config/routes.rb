# frozen_string_literal: true

Rails.application.routes.draw do
  ActiveAdmin.routes(self)

  scope '/(:currency)' do
    resources :items, only: %i[create edit update destroy]

    resources :consolidations, only: :index

    namespace :frontend do
      resources :cashes, only: %i[index edit update]

      namespace :dashboard do
        resources :cashes, only: %i[edit update index]
      end
    end

    get '(/:month)(/:category_id)', to: 'items#index', as: :root
  end
end
