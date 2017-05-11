RSpec.shared_examples :index do
  before { @format ||= :html }

  describe "#index.#{ @format }" do
    before { get :index, xhr: (@format == :js), format: @format }

    it { should render_template :index }
  end
end

RSpec.shared_examples :new do
  before { @format ||= :js }

  describe "#new.#{ @format }" do
    before { get :new, xhr: (@format == :js),  format: @format }

    it { should render_template :new }
  end
end

RSpec.shared_examples :create do
  before { @format ||= :js }

  before { expect(subject).to receive(:build_resource) }

  before { allow(subject).to receive(:resource).and_return(resource) }

  context do
    before { expect(resource).to receive(:save).and_return(true) }

    before { post :create, params: {}, format: @format }

    it { success.call }
  end

  context do
    before { expect(resource).to receive(:save).and_return(false) }

    before { post :create, params: {}, format: @format }

    it { failure.call }
  end
end

RSpec.shared_examples :show do
  before { @format ||= :html }

  describe "#show.#{ @format }" do
    before { get :show, xhr: (@format == :js), params: { id: 1 }, format: @format }

    it { should render_template :show }
  end
end

RSpec.shared_examples :edit do
  before { @format ||= :js }

  describe "#edit.#{ @format }" do
    before { get :edit, xhr: (@format == :js), params: { id: 1 }, format: @format }

    it { should render_template :edit }
  end
end

RSpec.shared_examples :update do
  before { @format ||= :js }

  before { allow(subject).to receive(:resource).and_return(resource) }

  before { expect(subject).to receive(:resource_params).and_return(:resource_params) }

  context do
    before { expect(resource).to receive(:update).with(:resource_params).and_return(true) }

    before { patch :update, params: { id: 1 }, format: @format }

    it { success.call }
  end

  context do
    before { expect(resource).to receive(:update).with(:resource_params).and_return(false) }

    before { patch :update, params: { id: 1 }, format: @format }

    it { failure.call }
  end
end

RSpec.shared_examples :destroy do
  let(:resource) { double }

  before { @format ||= :js }

  before { expect(subject).to receive(:resource).and_return(resource) }

  describe "#destroy.#{ @format }" do
    before { expect(resource).to receive(:destroy) }

    before { delete :destroy, params: { id: 1 }, format: @format }

    it { success.call }
  end
end
