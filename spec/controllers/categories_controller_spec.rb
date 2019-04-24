require 'rails_helper'

RSpec.describe CategoriesController, type: :controller do
  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { should eq :collection }
    end

    context do
      before { expect(Category).to receive(:order).with(:income).and_return(:collection) }

      its(:collection) { should eq :collection }
    end
  end

  describe '#resource' do
    context do
      before { subject.instance_variable_set :@resource, :resource }

      its(:resource) { should eq :resource }
    end

    context do
      before { expect(subject).to receive(:params).and_return({ id: 27 }) }

      before { expect(Category).to receive(:find).with(27).and_return(:resource) }

      its(:resource) { should eq :resource }
    end
  end

  it_behaves_like :edit

  describe '#resource_params' do
    let :params do
      acp category: { name: nil, income: nil, visible: nil }
    end

    before { expect(subject).to receive(:params).and_return(params) }

    its(:resource_params) { should eq params[:category].permit! }
  end

  it_behaves_like :update do
    let(:success) { -> { should render_template(:update).with_status(200) } }

    let(:failure) { -> { should render_template(:edit).with_status(422) } }
  end

  it_behaves_like :new

  it_behaves_like :create do
    let(:success) { -> { should render_template(:create).with_status(201) } }

    let(:failure) { -> { should render_template(:new).with_status(422) } }
  end
end
