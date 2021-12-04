# frozen_string_literal: true

RSpec.describe 'Backoffice::Categories', type: :request do
  pending do
    it_behaves_like 'new', '/backoffice/categories/new.js'
  end

  it_behaves_like 'edit', '/backoffice/categories/54/edit.js' do
    before { create :category, id: 54 }
  end

  pending do
    it_behaves_like 'create', '/backoffice/categories.js' do
      let :valid_params do
        {
          category: {
            name: 'Drinks',
            income: true,
            visible: true,
          }
        }
      end

      let :invalid_params do
        { category: { name: '' } }
      end

      let(:success) { -> { should render_template :create } }

      let(:failure) { -> { should render_template :new } }
    end
  end

  it_behaves_like 'update', '/backoffice/categories/54.js' do
    let :valid_params do
      { category: { name: 'Drinks' } }
    end

    let :invalid_params do
      { category: { name: '' } }
    end

    before { create :category, id: 54 }

    let(:success) { -> { should render_template :update } }

    let(:failure) { -> { should render_template :edit } }
  end

  describe 'GET /index.js' do
    before { create_list :category, 2 }

    before { get '/backoffice/categories', xhr: true }

    it_behaves_like 'index.js'
  end
end
