# frozen_string_literal: true

RSpec.describe 'Items', type: :request do
  it_behaves_like 'index', '/items'

  it_behaves_like 'index', '/items.js'

  it_behaves_like 'destroy', '/items/47.js' do
    before { create :item, id: 47 }

    let(:success) { -> { should render_template :destroy } }
  end

  pending do
    it_behaves_like 'create', '/items.js' do
      before { create :category, id: 13, name: 'Drinks' }

      let :valid_params do
        {
          item: {
            date: '2019-11-13',
            formula: '3 + 5',
            category_id: 13,
            description: 'Lorem Ipsum ...',
          }
        }
      end

      let :invalid_params do
        { item: { date: '' } }
      end

      let(:success) { -> { should render_template :create } }

      let(:failure) { -> { should render_template :new } }
    end
  end

  it_behaves_like 'edit', '/items/47/edit.js' do
    before { create :item, id: 47 }
  end

  it_behaves_like 'update', '/items/47.js' do
    let :valid_params do
      { item: { date: '2019-11-13' } }
    end

    let :invalid_params do
      { item: { date: '' } }
    end

    before { create :item, id: 47 }

    let(:success) { -> { should render_template :update } }

    let(:failure) { -> { should render_template :edit } }
  end
end
