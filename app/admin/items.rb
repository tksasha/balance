# frozen_string_literal: true

ActiveAdmin.register Item do
  include HasCurrencyScopes

  decorate_with ItemDecorator

  menu priority: 2, label: proc { I18n.t('active_admin.items') }

  actions :all, except: %i[destroy new create]

  permit_params %i[date]

  includes :category

  filter :category, as: :select, collection: []

  index row_class: proc { 'item' } do
    column :date
    column :sum, class: 'sum' do |item|
      money(item.sum)
    end
    column :category
    column :description do |item|
      sanitize(item.description)
    end
    actions
  end
end
