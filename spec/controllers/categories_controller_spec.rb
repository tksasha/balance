require 'rails_helper'

RSpec.describe CategoriesController, type: :controller do

  describe '#index' do
    before { get :index }

    it { should render_template :index }
  end

  describe '#show' do
    before { expect(Category).to receive(:find) }
    before { get :show, params: { id: 1 } }

    it { should render_template :show }
  end

  describe '#new' do
    before { get :new }

    it { should render_template :new }
  end

  describe '#create' do
    let(:category) { stub_model Category, name: 'Test', income: '0', slug: '', visible: '1' }
    let(:params) { { category: { name: 'Test', income: '0', slug: '', visible: '1' } } }

    before { expect(Category).to receive(:new).with(permit! params[:category]).and_return(category) }

    context do
      before { expect(category).to receive(:save).and_return(true) }
      before { post :create, params: params }

      it { should redirect_to category }
    end

    context do
      before { expect(category).to receive(:save).and_return(false) }
      before { post :create, params: params }

      it { should render_template :new }
    end
  end

  describe '#destroy' do
    let(:category) { stub_model Category }

    before { subject.instance_variable_set :@category, category }
    before { expect(category).to receive(:destroy) }

    before { delete :destroy, params: { id: 1 } }

    it { should redirect_to :categories }
  end

  describe '#edit' do
    before { expect(Category).to receive(:find) }
    before { get :edit, params: { id: 1 } }

    it { should render_template :edit }
  end

  describe '#update' do
    let(:category) { stub_model Category }
    let(:params) { { category: { name: 'Test', income: '0', slug: '', visible: '1' }, id: 1 } }

    before { subject.instance_variable_set :@category, category }

    context do
      before { expect(category).to receive(:update).with(permit! params[:category]).and_return(true) }
      before { put :update, params: params }

      it { should redirect_to category }
    end

    context do
      before { expect(category).to receive(:update).with(permit! params[:category]).and_return(false) }
      before { put :update, params: params }

      it { should render_template :edit }
    end
  end
end