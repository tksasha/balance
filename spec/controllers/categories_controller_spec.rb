# frozen_string_literal: true

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
      before { expect(subject).to receive(:params).and_return(id: 27) }

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

  describe '#initialize_resource' do
    before { expect(Category).to receive(:new).and_return(:resource) }

    before { subject.send :initialize_resource }

    its(:resource) { should eq :resource }
  end

  it_behaves_like :new

  describe '#build_resource' do
    before { expect(subject).to receive(:resource_params).and_return(:resource_params) }

    before { expect(Category).to receive(:new).with(:resource_params).and_return(:resource) }

    before { subject.send :build_resource }

    its(:resource) { should eq :resource }
  end

  it_behaves_like :create do
    let(:success) { -> { should render_template(:create).with_status(201) } }

    let(:failure) { -> { should render_template(:new).with_status(422) } }
  end

  describe '#set_variant' do
    context do
      before { expect(subject).to receive(:params).and_return({}) }

      it { expect(subject).to_not receive(:request) }

      after { subject.send :set_variant }
    end

    context do
      before { expect(subject).to receive(:params).and_return(widget: '') }

      it { expect(subject).to_not receive(:request) }

      after { subject.send :set_variant }
    end

    context do
      before { expect(subject).to receive(:params).and_return(widget: '1') }

      it { expect(subject).to receive_message_chain('request.variant=').with(:widget) }

      after { subject.send :set_variant }
    end
  end
end
